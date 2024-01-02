# EKS Networking

## Basics

### Control plane components

- Kube-apiserver: Exposes the K8s api. It's a frontend for the K8s control plane
- etcd: Key value store for K8s. Used as K8s backing store for all cluster data.
- kube-scheduler: watches for new pods with no node assigned and selects a node for them to run on
- kube-controller manager: runs controller processes such as node controller (if a node goes down, node controller will take care of this), replication controller, namespace controller, job controller, EndpointSlide controller etc
- cloud-controller manager: links K8s cluster into cloud provider's API such as node controller if node (instance) is deleted in the cloud, service controller for cloud load balancers

### Data plane/worker plane components

- Nodes: 
  - hosts the pods (containers)
- kubelet: an agent that runs on each node in the cluster. It makes sure that containers are running in a pod. Kubelet also register the state of your pods to the API server. 
- kube-proxy: enables network communication to pods from network sessions inside or outside of your cluster. 
- Container Runtime: responsible for running containers. K8s supports containers runtimes such as containerd, CRI-O, and any other implementation of the Kubernetes CRI (Container Runtime Interface).
  
### EKS Architecture

- K8s control plane and data plane are deployed across multiple AZs (min 3 AZs for HA and resiliency). EKS manages all the control plane components for you and you get a managed K8s control plane.
- When it comes to the data plane, you can either use self managed nodes (EC2) or managed node groups (EC2) or Fargate (no machines to manage).


## EKS Networking

### Cluster networking

- EKS has the control plane which is inside a managed VPC. 
- The worker nodes are launched in the customer defined VPC. 
- The control plane communicates with worker nodes via ENI which are injected into the customer VPC. SGs are attached to those ENIs so you can manage security. 
- EKS provisions 2-4 ENIs in the customer VPC to enable the communication between the control plane and the VPC (worker nodes). 
- It is advised or recommended to have separate subnets for EKS ENIs. EKS needs at least 6 IPs in each subnet (16 recommended) so /28 subnet. 
- EKS creates and associates SGs to the EKS owned ENIs (and also to manage groupd nodes)
- K8s API server can be accessed over the internet (by default)
- EKS allows you to assign IPv4 or IPv6 addresses to your pods. (but not in dual stack mode meaning pods can't have both IPv4 and IPv6 addresses at the same time)

### EKS Cluster endpoint access (pubic)

- EKS cluster endpoint public access is enabled by default and private access is disabled by default.
  - EKS will provide you the DNS which resolves to the public IP address of the K8s API server endpoint. 
- Worker nodes needs to be in a public subnet with an IGW so your control plane can communicate with the worker nodes.
- This means anyone can connect to your control plane. But obviously this is not secure. 
- You can restrict access to API public endpoing using CIDR blocks (to whitelist the IPs)

Private option:

- There's an option to enable both public and private accces for the control plane endpoint. In this case, the client can connect to the control plane over the internet, but the data plane can be in a private subnet and there is no need for an IGW.
- The control plane can access the data plane via the EKS owned ENI
- In this case, you can still restrict access to the API endpoint using CIDR blocks.
- This might be better than the first option but you still don't want even want the public connectivity to your control plane.
- You want all the connection from your network to the K8s control plane to be private. That means it's secured and over a VPN or a DX connection. You can completely disable public access to the control plane endpoint.
- In this case, you need some kind of Layer 4 connectivity between your network and the customer VPC. Typically this can be a VPN or a DX connection.
  - the client inside your network can reach out to the EKS owned ENI and then the traffic can go to your control plane.


In some cases, you want the client inside the private network to create the cluster. 

- in this case, you can use VPC interface endpoint for EKS. In the customer VPC, you create a VPC interface endpoint. 
- So the client will call the EKS API using eksctl or any other tool. The traffic will flow from the customer network to the customer VPC via a VPN or direct connect and then from there through to the VPC interface endpoint, and from there it can reach to the EKS service endpoint. 

Some additional best practices for EKS VPC external connectivity:

- Inside your customer VPC, have public subnets so that you can launch ELBs (ALB/NLB) so you can distribute traffic to your worker pods.
- Provide internet access to nodes/pods using NAT GW (IPv4) or egress-only IGW (IPv6) >> 
  - Egress only NGW only allows traffic from IPv6 addresses to go out to the internet and not in. 
  - But for IPv4 addresses, you should have NAT GWs. The NAT GW should be in public subnet, so traffic goes from your pods to the NAT GWs to the IGW to the internet.
- 
