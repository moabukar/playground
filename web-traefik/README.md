# Web app with traefik

## Build apps locally

- `docker build -f api.Dockerfile -t api-webapp .`
- `docker build -f frontend.Dockerfile -t frontend-webapp .` 

Create a docker network `docker network create traefik_webgateway`

## Start Traefik and app


- `docker-compose up`
- Access the app endpoint `curl http://mysuperapp.localhost/` or access it on the browser
- Access the app API endpoint `curl http://api.mysuperapp.localhost/docs` or access it on the browser
- Access Traefik dashboard `http://localhost:8080/dashboard/`
