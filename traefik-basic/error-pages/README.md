# Custom error pages image

- `docker build . -t error -f Dockerfile.error`
- `docker run -d -p 8095:80 error`

If you want to pull it from dockerhub:

- `docker run -d -p 8095:80 moabukar/error-pages:fcaa3c4`
