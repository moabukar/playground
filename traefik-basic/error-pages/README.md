# Custom error pages image

- `docker build . -t error -f Dockerfile.error`
- `docker run -d -p 8095:80 error`

## Directly from Docker Hub

- `docker run -d -p 8095:80 moabukar/error-pages:fcaa3c4`
