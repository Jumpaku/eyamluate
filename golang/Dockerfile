FROM ubuntu:23.10

RUN DEBIAN_FRONTEND=noninteractive apt update \
    && DEBIAN_FRONTEND=noninteractive apt install -y \
        ca-certificates \
        golang-go
ENV GOBIN=/usr/local/bin
RUN go install github.com/bufbuild/buf/cmd/buf@v1.39.0 \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2



WORKDIR /workspace
