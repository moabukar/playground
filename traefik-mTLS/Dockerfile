FROM ubuntu:22.04

ENV TRAEFIK_VERSION=2.10.3
ARG ARCH='linux_386'
ENV AWS_DEFAULT_REGION='eu-west-1'

RUN cd /tmp && \
    apt-get update && \
    apt install -y wget tar curl && \
    wget https://github.com/traefik/traefik/releases/download/v${TRAEFIK_VERSION}/traefik_v${TRAEFIK_VERSION}_${ARCH}.tar.gz && \
    tar -zxvf traefik_v${TRAEFIK_VERSION}_${ARCH}.tar.gz && \
    mv traefik /usr/bin/traefik && \
    chmod 0777 /usr/bin/traefik && \
    rm traefik_v${TRAEFIK_VERSION}_${ARCH}.tar.gz && \
    rm -rf /var/lib/apt/lists/*

RUN mkdir -p /etc/traefik && \
    mkdir -p /etc/traefik/conf && \
    mkdir -p /tls \
    mkdir -p /.step/authorities/test.ca/certs/

# Copy custom Traefik config files
COPY ./dynamic.yml /etc/traefik/conf/

COPY ./root_ca.crt /home/.step/authorities/test.ca/certs/root_ca.crt

WORKDIR /etc/traefik

# Expose ports
EXPOSE 80
EXPOSE 8080
EXPOSE 443

# Start Traefik
CMD ["traefik"]
