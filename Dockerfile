# CONTAINER FOR BUILDING BINARY
FROM golang:1.19 AS build

# INSTALL DEPENDENCIES
RUN go install github.com/gobuffalo/packr/v2/packr2@v2.8.3
COPY go.mod go.sum /src/
RUN cd /src && go mod download

# BUILD BINARY
COPY . /src
RUN cd /src/db && packr2
RUN cd /src && make build

# CONTAINER FOR RUNNING BINARY
FROM ubuntu:20.04
COPY --from=build /src/dist/cdk-data-availability /app/cdk-data-availability

# 添加Node.js源
RUN apt update && \
    apt install curl -y && \
    curl -sL https://deb.nodesource.com/setup_20.x | bash - && \
    apt-get install nodejs -y && \
    apt-get install python -y && \
    apt-get install make -y && \
    apt-get install vim -y && \
    npm install -g near-cli && \
    export NEAR_TESTNET_RPC=https://rpc.testnet.near.org && \
    near add-credentials example-acct.testnet --seedPhrase "antique attitude say evolve ring arrive hollow auto wide bronze usual unfold"

EXPOSE 8444
CMD ["/bin/sh", "-c", "/app/cdk-data-availability run"]
