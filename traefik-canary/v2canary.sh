#!/bin/bash

# Starting containers with Traefik v2 labels
docker run -d --name app_v1 \
  --label "traefik.enable=true" \
  --label "traefik.http.services.app-v1-service.loadbalancer.server.port=80" \
  nginx:stable

docker run -d --name app_v2 \
  --label "traefik.enable=true" \
  --label "traefik.http.services.app-v2-service.loadbalancer.server.port=80" \
  nginx:stable-alpine
