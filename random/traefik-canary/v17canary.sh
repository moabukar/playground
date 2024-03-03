#!/bin/bash

# Launch Traefik v1.7 as the reverse proxy
docker run -d \
  --name my_traefik \
  -p 8080:80 \
  -v /var/run/docker.sock:/var/run/docker.sock \
  traefik:v1.7 \
    --docker \
    --docker.exposedbydefault=false

# Start your application containers with new weight distributions
docker run -d --name app_v1 \
  --label "traefik.enable=true" \
  --label "traefik.backend=myapp" \
  --label "traefik.frontend.rule=Host:myapp.local" \
  --label "traefik.weight=30" \
  nginx:stable

docker run -d --name app_v2 \
  --label "traefik.enable=true" \
  --label "traefik.backend=myapp" \
  --label "traefik.frontend.rule=Host:myapp.local" \
  --label "traefik.weight=70" \
  nginx:stable-alpine
