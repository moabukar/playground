# NAT

## NAT instances

- Normal EC2 instance running an AMI configured to provide NAT services
  - no longer recommened because it's running an Amazon Linux, 2018.03 AMI (end of life)
  - AWS recommends customers to move to NAT Gateway
  - non-specialised, just EC2 instance running NAT software
- Speed based on size and type of EC2 instance
- Disable source/dest check (if interface is neither the source nor destination then by default it will drop the packets) (so enable this option if you're using NAT instances)
- SGs on the NAT instance & NACL on the subnet
- Can also be used as a bastion host and can be configured for port forwarding
- ...self managed, so you have to patch it, monitor it, etc.
- Scripts are used to implement HA
  - they are just EC2 instances so failures need to be managed with thes scripts
  - route table changes based on failures
- Using NAT instances can help reduce costs as you can combine different functionalities
  - combine NAT instance with a bastion host and port forwards all running from the same EC2 instance
- SG & NACL can be used to control access to the NAT instance
- Flows logs for monitoring - to review packet metadata that's flowing through the VPC. You could put a VPC flow log to specifically monitor the network interface of the NAT instance or scope it wider by putting it on the subnet or VPC
- NAT instance can be used across DX, VPN, peer connections (not ideal but possible)
 
## NAT Gateway (AWS Recommended)

- 
