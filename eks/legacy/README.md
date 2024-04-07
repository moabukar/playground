# EKS

Implementation of EKS setup using `Terraform` and `Cloudformation`. Fully functional templates to deploy your `VPC` and `Kubernetes clusters` together with all the essential tags and addons. Also, worker nodes are part of AutoScallingGroup which consists of spot and on-demand instances.

Templates support deployment to different AWS partitions. I have tested it with `public` and `china` partitions. I am actively using this configuration to run EKS setup in Ireland(eu-west-1), North Virginia(us-east-1) and Beijing(cn-north-1).

### terraform-aws and terraform-k8s templates

Latest configuration templates used by me can be found in [terraform-aws](./terraform-aws/) for aws provider and [terraform-k8s](./terraform-k8s/) for kubernetes provider. Once you configure your environment variables in `./terraform-aws/vars` `./terraform-k8s/vars`, you can use makefile commands to run your deployments. Resources that will be created after applying templates:

You will find latest setup of following components:

1. VPC with public/private subnets, enabled flow logs and VPC endpoints for ECR and S3
1. EKS controlplane
1. EKS worker nodes in private subnets (spot and ondemnd instances based on variables)
1. Karpenter configuration for nodes
1. Option to used Managed Node Groups
1. Dynamic basion host
1. Automatically configure aws-auth configmap for worker nodes to join the cluster
1. OpenID Connect provider which can be used to assign IAM roles to service accounts in k8s
2. NodeDrainer lambda which will drain worker nodes during rollingUpdate of the nodes (This is only applicable to spot worker nodes, managed node groups do not require this lambda). 
3. IAM Roles for service accounts such as aws-node, cluster-autoscaler, alb-ingress-controller, external-secrets (Role arns are used when you deploy kubernetes addons with Service Accounts that make use of OIDC provider)
4. For spot termination handling use aws-node-termination-handler from [k8s_templates/aws-node-termination-handler](./k8s_templates/aws-node-termination-handler).
5. EKS cluster add-ons (CoreDNS + kube-proxy)

## Kubernetes YAML templates

All the templates for additional deployments/daemonsets can be found in [k8s_templates](./k8s_templates/).

To apply templates simply run `kubectl apply -f .` from a desired folder. Ensure to put in correct Role arn in service accounts configuration. Also, check that environment variables are correct. 

You will find templates for the following Kubernetes components:

* ALB ingress controller
* AWS Load Balancer controller
* AWS node termination handler
* Calico
* Cert Manager
* Cluster Autoscaler
* CoreDns
* Dashboard
* External-DNS
* External Secrets
* External Secrets Operator (helm chart using fluxv2)
* Karpenter (helm chart using fluxv2)
* Kube Proxy
* Kube2iam
* Metrics server
* NewRelic
* Reloader
* Spot Interrupt Handler
* Vertical Pod Autoscaler with cert-manager certificate
* VPC CNI Plugin
* Secrets CSI Driver

### Useful resources

[EKS platforms information](https://docs.aws.amazon.com/eks/latest/userguide/platform-versions.html)
[Worker nodes upgrades](https://docs.aws.amazon.com/eks/latest/userguide/update-stack.html)

## Generate kubeconfig file

On user's machine who has been added to EKS, they can configure .kube/config file using the following command:

```bash
$ aws eks list-clusters
$ aws eks update-kubeconfig --name ${cluster_name}
```
