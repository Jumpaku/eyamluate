FROM --platform=linux/amd64 dart:3.5.2

RUN curl -sSL https://github.com/protocolbuffers/protobuf/releases/download/v28.0/protoc-28.0-linux-x86_64.zip -o /tmp/protoc.zip \
    && unzip /tmp/protoc.zip -d /usr/local \
    && rm /tmp/protoc.zip
RUN curl -sSL https://github.com/bufbuild/buf/releases/download/v1.40.1/buf-Linux-x86_64 -o /tmp/buf \
    && chmod +x /tmp/buf \
    && mv /tmp/buf /usr/local/bin/buf
RUN dart pub global activate protoc_plugin


