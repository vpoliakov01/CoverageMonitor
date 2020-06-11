FROM ubuntu:20.04

ARG DEBIAN_FRONTEND=noninteractive

ENV GO_VERSION="1.13.4"

# Install prereqs
RUN apt-get update && \
    apt install -y git curl make build-essential nginx-full nodejs npm

# Copy the repo
RUN mkdir /CoverageMonitor
COPY back_end /CoverageMonitor/back_end
COPY front_end /CoverageMonitor/front_end
COPY nginx.conf /CoverageMonitor
COPY go.mod /CoverageMonitor
COPY go.sum /CoverageMonitor
COPY makefile /CoverageMonitor

# Setup React, TypeScript, and Webpack
RUN cd /CoverageMonitor/front_end  npm install --save react react-dom && \
    npm install --save-dev webpack webpack-cli @types/react @types/react-dom typescript ts-loader source-map-loader
RUN cd /CoverageMonitor && make client

# Install Go
RUN curl -L -o /tmp/go.tgz https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xvf /tmp/go.tgz && \
    rm -f /tmp/go.tgz && \
    mv /usr/local/go/bin/go /usr/local/bin/go
RUN cd /CoverageMonitor && make server

# Build the server
ENTRYPOINT [ "bash", "-c", "cd /CoverageMonitor && make app" ]
