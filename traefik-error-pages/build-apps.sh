#!/bin/bash

# Build the Docker image
docker build . -t nodeapp -f Dockerfile.node

# Start containers
docker run -d -p 9991:9999 -e CONTAINER_NUMBER=1 nodeapp
docker run -d -p 9992:9999 -e CONTAINER_NUMBER=2 nodeapp
docker run -d -p 9993:9999 -e CONTAINER_NUMBER=3 nodeapp
docker run -d -p 9994:9999 -e CONTAINER_NUMBER=4 nodeapp

echo "Containers are up and running!"

##
echo "Starting error containers served by NGINX"

docker build . -t nginx -f Dockerfile.error
docker run -d -p 8090:80 nginx

echo "all ready!"
## Add some lines below to start Traefik too?
