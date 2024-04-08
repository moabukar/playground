# Setup

```sh
docker build -t ghr .
docker run -e GITHUB_RUNNER_TOKEN=<TOKEN_HERE> ghr
```

## K8s 

```sh

echo -n 'token' | base64

```
