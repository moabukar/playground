# EC2 Networking

- When dealing with EC2 networking, the most important part is the ENI (Elastic Network Interface) - they control how networking is managed for EC2 instances. Security groups and NACLs are attached to ENIs, and ENIs are attached to EC2 instances. 
- ENIs can be moved between EC2 instances. An EC2 instance can have multiple ENIs attached to it at the same time.
- The ability to attach multiple ENIs to a single instance is useful for various scenarios, such as:
  - Creating dual-homed instances: Instances with workloads or roles on different subnets or networks.
  - Managing traffic: Separating different types of traffic for security or performance reasons.
  - High Availability setups: Using secondary interfaces for failover or redundancy.

