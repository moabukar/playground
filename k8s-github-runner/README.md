# K8s GitHub runner

## setup

```sh

## Create a GitHub Personal Access Token with the repo and admin:org scope permissions:

## Create a generic secret with your GitHub Personal Access Token

kubectl create secret generic my-pat --from-literal=pat=XXXXXXXXXXXXXXXXXX

### If you want to include your registry self-signed CA certificate, use a config map

kubectl create configmap private-registry-certificate --from-file=ca.crt

## add GITHUB_OWNER environment variable in your deploy, and deploy:
kubectl create -f runner-k8s/deploy.yml

## you should see your runners now!
