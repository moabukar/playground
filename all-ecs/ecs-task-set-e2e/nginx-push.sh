VERSION=$1

set -o errexit

docker build --tag AWS_ACCOUNT_ID.dkr.ecr.eu-west-1.amazonaws.com/mo-nginx:${VERSION} .
aws ecr get-login-password --region eu-west-1 | docker login --username AWS --password-stdin AWS_ACCOUNT_ID.dkr.ecr.eu-west-1.amazonaws.com
docker push AWS_ACCOUNT_ID.dkr.ecr.eu-west-1.amazonaws.com/mo-nginx:${VERSION}

#nginx:stable-alpine3.17-slim
