# [Choice] Go version: 1, 1.15, 1.14
ARG VARIANT=1.14
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# [Optional] Uncomment this section to install additional OS packages.
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends acl musl-tools snapcraft

# [Optional] Uncomment the next line to use go get to install anything else you need
RUN go get -x mvdan.cc/gofumpt

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

RUN curl -sfL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | sh -s -- -b $(go env GOPATH)/bin
