FROM golang:1.16.3-alpine3.13 as builder

# System setup
RUN apk update && apk add git curl build-base autoconf automake libtool

# Install protoc
ENV PROTOBUF_URL https://github.com/google/protobuf/releases/download/v3.3.0/protobuf-cpp-3.3.0.tar.gz
RUN curl -L -o /tmp/protobuf.tar.gz $PROTOBUF_URL
WORKDIR /tmp/
RUN tar xvzf protobuf.tar.gz
WORKDIR /tmp/protobuf-3.3.0
RUN mkdir /export
RUN ./autogen.sh && \
    ./configure --prefix=/export && \
    make -j 3 && \
    make check && \
    make install

# Install protoc-gen-go
WORKDIR /golang
COPY tools.go .
COPY go.mod .
# COPY go.sum .
#RUN go 
RUN go mod tidy
# RUN go get \
#     github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
#     github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
#     github.com/golang/protobuf/protoc-gen-go
RUN go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
#RUN go get github.com/golang/protobuf/protoc-gen-go
RUN cp /go/bin/protoc-gen-go-grpc /export/bin/
RUN cp /go/bin/protoc-gen-openapiv2 /export/bin/
RUN cp /go/bin/protoc-gen-grpc-gateway /export/bin/
RUN cp /go/bin/protoc-gen-go /export/bin/
# Export dependencies
RUN cp /usr/lib/libstdc++* /export/lib/
RUN cp /usr/lib/libgcc_s* /export/lib/
WORKDIR /tmp     
RUN git clone https://github.com/googleapis/googleapis 

FROM golang:1.16.3-alpine3.13
RUN apk update && \
    apk add --no-cache tzdata ca-certificates alpine-conf libstdc++ clang
RUN /sbin/setup-timezone -z Europe/Amsterdam

COPY --from=builder /export /usr
COPY --from=builder /tmp/googleapis/google /app/google
WORKDIR /app
RUN ls -la



