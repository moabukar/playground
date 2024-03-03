# Custom error pages image

- `docker build . -t error -f Dockerfile`
- `docker run -d -p 8095:82 error`

## Directly from Docker Hub

- `docker run -d -p 8095:82 moabukar/error-pages:fcaa3c4`
