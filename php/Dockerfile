FROM --platform=linux/amd64 composer:2.2.24 AS composer

FROM --platform=linux/amd64 php:8.3 AS workspace

COPY --from=composer /usr/bin/composer /usr/bin/composer

RUN DEBIAN_FRONTEND=noninteractive apt update \
    &&  DEBIAN_FRONTEND=noninteractive apt install -y unzip

RUN curl -sSL https://github.com/protocolbuffers/protobuf/releases/download/v28.0/protoc-28.0-linux-x86_64.zip -o /tmp/protoc.zip \
    && unzip /tmp/protoc.zip -d /usr/local \
    && rm /tmp/protoc.zip
RUN curl -sSL https://github.com/bufbuild/buf/releases/download/v1.40.1/buf-Linux-x86_64 -o /tmp/buf \
    && chmod +x /tmp/buf \
    && mv /tmp/buf /usr/local/bin/buf
