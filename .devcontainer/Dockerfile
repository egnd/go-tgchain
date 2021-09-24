ARG VARIANT="1.17"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

ENV TZ Europe/Moscow
ENV CGO_ENABLED 0
ENV GOPROXY https://proxy.golang.org,direct
ENV GOSUMDB off

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive && \
    apt-get -y install --no-install-recommends git make grep && \
    which git make grep
