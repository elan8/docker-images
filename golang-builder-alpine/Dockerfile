FROM golang:1.16.3-alpine3.13 as builder

# System setup
RUN apk update && apk add build-base git autoconf automake libtool curl make g++ unzip tar
RUN gcc --version

# Install protoc
ENV PROTOBUF_URL https://github.com/google/protobuf/releases/download/v3.16.0/protobuf-cpp-3.16.0.tar.gz
ENV CXXFLAGS="-std=c++11"
RUN curl -L -o /tmp/protobuf.tar.gz $PROTOBUF_URL
WORKDIR /tmp/
RUN tar xvzf protobuf.tar.gz >> /dev/null
WORKDIR /tmp/protobuf-3.16.0
RUN mkdir /export
RUN ./autogen.sh && \
    ./configure --prefix=/export 
RUN make && make install


WORKDIR /golang
COPY tools.go .
COPY go.mod .
RUN go mod tidy
COPY json-pg json-pg
RUN ls -la
RUN cd json-pg && go install
RUN go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc \
    github.com/favadi/protoc-go-inject-tag
RUN cd /go/bin/ && ls -la
RUN cp /go/bin/json-pg /export/bin/
RUN cp /go/bin/protoc-go-inject-tag /export/bin/
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



