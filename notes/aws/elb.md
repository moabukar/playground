# ELB

- ELB nodes have public IP & private IPs if they are internet facing.
- ELB nodes have only private IPs if they are internal.
- Each ALB is configured with an A records DNS name. This resolves to the public IP addresses of the ELB nodes

- Load balancer nodes are configured with listeners which accept traffic on a port & protocol and communicate with targets on a port & protocol.
- Only requirement is that load balancer nodes can communicate with the backend instances 
- LBs in order to function, need 8 or more free IP addresses in the subnets. So strictly speaking, this means a /28 subnet which gives you around 16 IPs (minus the 5 reserved by AWS) which leaves 11 free per subnet.
    - but AWS suggests you use a /27 or larger subnet to deploy an ELB in order for it to scale.
- Internal LBs are architecturally just like internet-facing LBs except they only have private IPs allocated to their nodes. So, internal LBs are generally used to separate differnt tiers of apps
- So here, User A connects via the internet-facing LB to the web server. And then this web server can connect to an app server via an internal LB. This allows us to separate app tiers and allow for independent scaling of each tier.

----------------
## Example of a web, app and a DB architecture
  - User connects to internet-facing LB which connects to the web server. The web server connects to the app server via an internal LB. The app server connects to the DB directly. 
  - Note: you have a separate ASG for each the web and app ASG. 
- LB's allow each tier to scale independently as there is no loose coupling. 

## Cross Zone Load Balancing

- User A accesses a wordpress blog. So User A uses his device and browses the app which is actually the DNS name.
- The LB DNS name will load balance the traffic across 2 LB nodes in different AZs. Each node will receive a portion of the incoming requests based on how many nodes there are. 
  - For 2 nodes, each node gets 50% of the traffic. (100% / number of nodes)
  - So if you had 4 instances in LB node 1, each instance will get 25% of the traffic. and if LB node 2 had only 1 instance, that instance will get 100% of the traffic.
  - So the traffic is distributed across the nodes, but the nodes are not distributed across the AZs. So these are unevenly distributed which can cause issues
  - The fix for this was cross-zone load balancing.
- Cross-zone load balancing simply allows every load balancer node to distribute any connections that it receieves equally across all registered instances in all AZs. 
- So this means, the node in AZ 1 could distributed connections to the intance in AZ 2 and vice versa. And this represents a more even distribution of traffic across the AZs.
- This feature was originally not enabled by default. But these days if you are deploying an ALB, this comes enabled as a standard.

### Cross Zone LB Architecture & Summary

- When you create an ELB, you see it as one device which runs in two or more AZs. One subnet in each AZ. But what you're actually creating is one ELB node in one subnet in each AZ. 
- You're also creating a DNS A record for that LB which spreads the incoming requests over all the active nodes for that LB. 
- One node (one subnet per AZ) can scale automatically if additonal load if placed on the LB.
- Remember by default, cross zone load balancing means that nodes can distribute requests across to other AZs. But historically this was disabled meaning connections potentially could be distributed unevenly across AZs. But for ALBs now, cross zone load balancing is enabled by default.
- LBs come in 2 flavours: 
  - internet-facing - nodes have public IPv4s 
  - internal - nodes have private IPs
- Important note: EC2 doesn't need to be public to communicate with an LB. The LB can be internet-facing but the EC2 instances can be private and don't need public IP addressing. 
- An internet-facing LB has public IP addresses on its nodes. It can accept connection from the public internet and distribute these across both public and private EC2 instances.
- LBs are configured via listener configurations and it controls WHAT the LBs listen to.
- LBs require 8+ free IPs per subnet and a /27 or larger subnet to allow for scaling. Strictly speaking, you need a /28 subnet would be enough but AWS suggests a /27 subnet. 

## User Session State

- Server-side piece of information
- Persists while you interact with the app
- Shopping cart, workflow position or login state
- Session state loss = Bad user user experience (UX) or issues
- Session state stored on a server or externally. If session state is stored externally, then those servers are stateless.

### Why session state matters:

- User A is browsing amazon.com. Imagine the app is running through an ELB. 
  - The ELB's job is to abstract User A from the underlying infrastructure.
  - And the 2nd function of the LB is to distribute the request across all the registered compute.

- Assume user A connects to instance 2 and assume user A browses the site and adds items to the cart. But he needs to get his card to order. So user A's session data is stored on instance 2. Let's say instance 2 fails, the LB is smart enough to immediately re-route the connection to instance 1. So instance 1 has a different set of user session data and so user A loses his cart and potentially gets logged out. 
- 
