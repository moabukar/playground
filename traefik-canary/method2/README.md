# Step by step

Step 2: Run Traefik Container with Docker and File Providers
Now you're going to run the Traefik container and mount both the Docker socket (so it can interact with Docker) and your static.yml.

```bash

docker run -d \
  --name traefik-v2.5 \
  -p 80:80 \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v $(pwd)/static.yml:/etc/traefik/static.yml \
  traefik:v2.5 \
  --providers.docker \
  --providers.docker.exposedbydefault=false \
  --providers.file.filename=/etc/traefik/static.yml


```

## Start apps

```bash

# Run the stable version of your app
docker run -d \
  --name app_normal_01 \
  --label "traefik.enable=true" \
  --label "traefik.http.services.app_normal.loadbalancer.server.port=80" \
  nginx:1.19.1

# Run the canary (new version) of your app
docker run -d \
  --name app_canary_01 \
  --label "traefik.enable=true" \
  --label "traefik.http.services.app_canary.loadbalancer.server.port=80" \
  nginx:1.19.2


```


## Test

```bash

for i in {1..100}; do 
  curl -s -o /dev/null -D - -H "Host:example.local" "http://localhost" | grep "Server"; 
done | sort | uniq -c


```
