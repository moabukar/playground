# NGINX setup

## Local setup

1. `cd app`
2. `docker build -t nginx .`
3. `docker run --rm --name nginx -p 80:80 -d nginx`
