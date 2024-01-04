VERSION=$1

set -o errexit

docker build --tag AWS_ACCOUNT_ID.dkr.ecr.eu-west-1.amazonaws.com/mo-apache:${VERSION} .
aws ecr get-login-password --region eu-west-1 | docker login --username AWS --password-stdin AWS_ACCOUNT_ID.dkr.ecr.eu-west-1.amazonaws.com
docker push AWS_ACCOUNT_ID.dkr.ecr.eu-west-1.amazonaws.com/mo-apache:${VERSION}

#httpd:alpine3.19
