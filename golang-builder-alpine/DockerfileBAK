FROM golang:1.15-alpine as builder 
WORKDIR /tmp
# System setup
RUN apk update && \
    apk add --no-cache tzdata ca-certificates alpine-conf git curl build-base autoconf automake libtool clang g++ unzip make cmake
RUN /sbin/setup-timezone -z Europe/Amsterdam
# build and install protoc
ENV PROTOBUF_URL https://github.com/google/protobuf/releases/download/v3.14.0/protobuf-cpp-3.14.0.tar.gz
RUN curl -L -o /tmp/protobuf.tar.gz $PROTOBUF_URL && \
    tar xvzf protobuf.tar.gz && \
    cd /tmp/protobuf-3.14.0 && \
    mkdir /export && \
    cd cmake && cmake -Dprotobuf_BUILD_TESTS=OFF -DCMAKE_INSTALL_PREFIX=/export && \
    make -j 3 && \
    make install && rm -R /tmp/*
WORKDIR /tmp     
RUN git clone https://github.com/googleapis/googleapis   

FROM golang:1.15-alpine 
RUN apk update && \
    apk add --no-cache tzdata ca-certificates alpine-conf libstdc++ clang
RUN /sbin/setup-timezone -z Europe/Amsterdam
WORKDIR /golang
COPY tools.go .
COPY go.mod .
COPY go.sum .
RUN go install \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
    github.com/golang/protobuf/protoc-gen-go
COPY --from=builder /export/lib64 /usr/lib
COPY --from=builder /export/bin /usr/bin
COPY --from=builder /export/include/google /usr/include/google
COPY --from=builder /tmp/googleapis/google /app/google
WORKDIR /app


