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

- 
