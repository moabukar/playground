# Traefik ECS Tasksets

## Using Cloudmap & Service Discovery

- Nginx >> http://nginx.mo-sandbox.sandbox.<hosted_zone>.services/
- Apache >> http://apache.mo-sandbox.sandbox.<hosted_zone>.services/

- Route for both >> http://both.mo-sandbox.sandbox.<hosted_zone>.services/

## Using Traefik for the Task sets

- Once all ECS infra is setup and task sets are running, run steps below to test:

```bash

sudo rm -rf /etc/traefik/conf/dynamic.yml

sudo ln -s $(pwd)/dynamic.yml /etc/traefik/conf/dynamic.yml

add this to your /etc/hosts file:
127.0.0.1 localhost both.mo-sandbox.sandbox.<hosted_zone>.services

create a DNS entry on cloud map

traefik --configfile=static.yml

http://localhost:8080/dashboard/#/

browse or curl:

http://both.mo-sandbox.sandbox.<hosted_zone>.services/

This should do weighted routing between NGINX & Apache services deployed on ECS 

```
