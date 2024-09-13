FROM --platform=linux/amd64 node:22-bookworm-slim

RUN DEBIAN_FRONTEND=noninteractive apt update \
    &&  DEBIAN_FRONTEND=noninteractive apt install -y curl unzip

RUN curl -sSL https://github.com/protocolbuffers/protobuf/releases/download/v28.0/protoc-28.0-linux-x86_64.zip -o /tmp/protoc.zip \
    && unzip /tmp/protoc.zip -d /usr/local \
    && rm /tmp/protoc.zip
RUN curl -sSL https://github.com/bufbuild/buf/releases/download/v1.40.1/buf-Linux-x86_64 -o /tmp/buf \
    && chmod +x /tmp/buf \
    && mv /tmp/buf /usr/local/bin/buf

RUN npm install -g typescript@5.6.2 rollup@4.21.2 terser@5.32.0 @swc/cli @swc/core
