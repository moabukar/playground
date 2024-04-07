# FROM debian:buster-slim
FROM --platform=linux/arm64/v8 debian:buster-slim

ARG GITHUB_RUNNER_VERSION="2.293.0"
# ARG HELM_VERSION="3.9.0"

ENV GITHUB_PAT ""
ENV GITHUB_OWNER "moabukar"
ENV RUNNER_WORKDIR "_work"
ENV RUNNER_LABELS ""
ENV RUNNER_NAME_PREFIX "myorg-"

SHELL ["/bin/bash", "-o", "pipefail", "-c"]

RUN apt-get update \
    && apt-get install --no-install-recommends -y \
        apt-transport-https \
        ca-certificates \
        gnupg \
        gnupg-agent \
        software-properties-common \
        curl \
        git \
        jq \
        curl \
        xvfb \
        libgtk2.0-0 \
        libgtk-3-0 \
        libgbm-dev \
        libnotify-dev \
        libgconf-2-4 \
        libnss3 \
        libxss1 \
        libasound2 \
        libxtst6 \
        xauth \
        jq \
        wget \
    && rm -rf /var/lib/apt/lists/*


# Install helm
# RUN wget https://get.helm.sh/helm-v${HELM_VERSION}-linux-amd64.tar.gz \
#     && tar -zxvf helm-v${HELM_VERSION}-linux-amd64.tar.gz \
#     && mv linux-amd64/helm /usr/local/bin/helm \
#     && rm helm-v${HELM_VERSION}-linux-amd64.tar.gz

# Install Docker client
RUN curl -fsSL https://download.docker.com/linux/debian/gpg | apt-key add - \
    && echo "deb [arch=$(dpkg --print-architecture)] https://download.docker.com/linux/debian $(lsb_release -cs) stable" > /etc/apt/sources.list.d/docker.list \
    && apt-get update \
    && apt-get install docker-ce-cli --no-install-recommends -y \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*
    # && useradd -m github

WORKDIR /home/github

# Install github runner packages
RUN curl -Ls https://github.com/actions/runner/releases/download/v${GITHUB_RUNNER_VERSION}/actions-runner-linux-x64-${GITHUB_RUNNER_VERSION}.tar.gz | tar xz \
    && ./bin/installdependencies.sh 
    # && chown -R github:github /home/github

# USER github

# Entrypoint
# COPY --chown=github:github entrypoint.sh ./entrypoint.sh
COPY entrypoint.sh ./entrypoint.sh
RUN chmod u+x ./entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]
