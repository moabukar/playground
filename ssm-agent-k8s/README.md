# K8s SSM agent

- A simple SSM agent to allow for secure access to K8s worker nodes.

## Setup

```sh
## Deploy daemonset to cluster
kubectl apply -f k8s/ds.yaml


## Access EC2 instance

aws --region eu-west-1 ssm start-session --target i-<instance-id>
```
