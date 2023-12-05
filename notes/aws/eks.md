# Commong challenges faced

## Run out of IP address spaces & latency

- Use prefix assignment
- Add secondary CIDR blocks
- Use overlapping non routable CIDR blocks (different regions?)
- Use IPv6 (but what are the challenges of these?)

## Cluster Autoscaler is slow during traffic spikes

- Use AWS Karpenter 
  - More performant than Cluster Autoscaler
  - Spins up in seconds compared to minutes of cluster autoscaler
  - it can directly access the AWS API to spin up instances


