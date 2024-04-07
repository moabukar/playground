FROM debian:buster-slim

RUN apt-get update && apt-get install -y \
    curl \
    sudo \
    git \
    jq \
    tar \
    gzip \
    ca-certificates \
    libicu63 \
    --no-install-recommends && \
    rm -rf /var/lib/apt/lists/*

ARG GITHUB_RUNNER_VERSION="2.315.0"
ARG RUNNER_ARCH="arm64"

RUN useradd -m runner && \
    echo "runner ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers

WORKDIR /home/runner

USER runner

# Download the GitHub Actions runner for Linux ARM64
RUN curl -o actions-runner-linux-${RUNNER_ARCH}-${GITHUB_RUNNER_VERSION}.tar.gz -L \
    https://github.com/actions/runner/releases/download/v${GITHUB_RUNNER_VERSION}/actions-runner-linux-${RUNNER_ARCH}-${GITHUB_RUNNER_VERSION}.tar.gz

# Validate the runner's checksum - replace with the checksum you have for your specific version
# Ensure you have the correct checksum for Linux ARM64 runner
RUN echo "d9d58b178eca5fb65d93d151f3b62bde967f8cbec7c72e9b0976e9312b7f7dda  actions-runner-linux-${RUNNER_ARCH}-${GITHUB_RUNNER_VERSION}.tar.gz" | shasum -a 256 -c

RUN tar xzf ./actions-runner-linux-${RUNNER_ARCH}-${GITHUB_RUNNER_VERSION}.tar.gz && \
    rm ./actions-runner-linux-${RUNNER_ARCH}-${GITHUB_RUNNER_VERSION}.tar.gz

COPY entrypoint.sh .
RUN sudo chmod +x ./entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]
