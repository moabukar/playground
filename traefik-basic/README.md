# Traefik - Load balancing between containers/apps

Without Docker (directly via Traefik CLI):

- Install Traefik (`brew install traefik`)
- `cd app`
- `ln -s $(pwd)/dynamic.yml /etc/traefik/conf/dynamic.yml`
- `traefik --configfile=./static.yml`
- `http://localhost:8080/dashboard`

With Docker:

- `docker build . -t traefik`
- `docker build . -t traefik --build-arg ARCH=linux_arm64` // for M1 Macs
- `docker run -p 8080:8080 -p 80:80 -v ./static.yml:/etc/traefik/traefik.yml -it traefik`
- `http://localhost:8080/dashboard`

## Using Traefik to load balance between apps

## Single build

- `docker build . -t nodeapp -f Dockerfile.node`
- `docker run --name nodeapp -p 9999:9999 nodeapp`

## Multiple builds (used with Traefik for LB)

- `docker build . -t nodeapp -f Dockerfile.node`
- `docker run -d -p 9991:9999 -e CONTAINER_NUMBER=1 nodeapp` - on port 9991
- `docker run -d -p 9992:9999 -e CONTAINER_NUMBER=2 nodeapp` - on port 9992
- `docker run -d -p 9993:9999 -e CONTAINER_NUMBER=3 nodeapp` - on port 8002
- `docker run -d -p 9994:9999 -e CONTAINER_NUMBER=4 nodeapp` - on port 8002

## Traefik build

- `traefik --configfile=./static.yml`
- `http://localhost:80` >> This will load balance between the multiple apps running on containers.
