# Traefik - Load balancing between containers/apps

Without Docker (directly via Traefik CLI):

- Install Traefik (`brew install traefik`)
- `cd app`
- `ln -s $(pwd)/dynamic.yml /etc/traefik/conf/dynamic.yml`
- `ln -s $(pwd)/static.yml /etc/traefik/traefik.yml`
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
- `docker run -d -p 9993:9999 -e CONTAINER_NUMBER=3 nodeapp` - on port 9993
- `docker run -d -p 9994:9999 -e CONTAINER_NUMBER=4 nodeapp` - on port 9994

## Traefik build

- `traefik --configfile=./static.yml`
- `http://localhost:80` >> This will load balance between the multiple apps running on containers.


## Using Jaeger tracing alongside Traefik

```bash
- docker run \
    -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p5775:5775/udp -p6831:6831/udp -p6832:6832/udp \
    -p5778:5778 -p16686:16686 -p14268:14268 -p9411:9411 jaegertracing/all-in-one:latest
```

- Enable Jaeger tracing in `static.yml`:

```bash
tracing:
  jaeger:
    samplingServerURL: http://localhost:5778/sampling
    localAgentHostPort: "localhost:6831"

```

- Once both above are setup, you can view it on the Jaeger UI here `http://localhost:16686/`


## Error pages served by custom error image (runs NGINX in the background)

- cd into `error-pages` directory
- `docker build . -t error -f Dockerfile.error`
- `docker run -d -p 8095:80 error` >> This will show the error page served by the custom error image (runs NGINX inside)

- Create a service in Traefik to route to the serror pages. Which will be used by the middleware to serve the error page.

- Make sure apps are not running first then browse `weighted-svc.localhost` and you should see the 502 error served by Traefik custom error page.

- Also test out `weighted-svc.localhost` on browser which should load normally. Then test out `weighted-svc.localhost/random/path` and you should see the error served by NGINX. 


## ALL DEMO:

- `docker compose up --build` or `docker compose up --build -d`
- Run Traefik `traefik --configfile=./static.yml`
- Test out `weighted-svc.localhost` on browser which should load normally. Then test out `weighted-svc.localhost/random/path` and you should see the error served by NGINX.

## Simulate errors (4xx and 5xx)

- Spam requests on weighted-svc.localhost and you should see a custom 429 error served by the error container. 
- Try access `http://weighted-svc.localhost/error-503` to see a custom 503 error served by the error container.

### All with docker compose (Ignore this)

- `docker network create traefik-net`
- `docker compose up --build -d`

### Others (Ignore this)

- `docker build . -t traefik -f Dockerfile.traefik`
- `docker build . -t traefik -f Dockerfile.traefik --build-arg ARCH=linux_arm64` // for M1 Macs
- `docker run -p 8080:8080 -p 80:80 -v ./static.yml:/etc/traefik/traefik.yml -it traefik`
- `http://localhost:8080/dashboard`


- docker run -p 8080:8080 -p 80:80 -v ./static.yml:/etc/traefik/traefik.yml -it traefik
