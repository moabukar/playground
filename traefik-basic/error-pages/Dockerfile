FROM alpine:3.12

ENV NGINX_VERSION 1.19.1
ENV NJS_VERSION 0.4.2
ENV PKG_RELEASE 1


RUN apk add --no-cache nginx curl libssl1.1 libcrypto1.1
COPY nginx.conf /etc/nginx/nginx.conf
COPY default.conf /etc/nginx/conf.d/default.conf
COPY ./pages /var/www/public


EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
