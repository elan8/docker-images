FROM golang:1.12-alpine3.9 
WORKDIR /app
# System setup
RUN apk update && apk add git curl build-base autoconf automake libtool clang
# build and install protoc
ENV PROTOBUF_URL https://github.com/google/protobuf/releases/download/v3.6.0/protobuf-cpp-3.6.0.tar.gz
RUN curl -L -o /tmp/protobuf.tar.gz $PROTOBUF_URL
WORKDIR /tmp/
RUN tar xvzf protobuf.tar.gz
WORKDIR /tmp/protobuf-3.6.0
RUN mkdir /export
RUN ./autogen.sh && \
    ./configure --prefix=/export && \
    make -j 3 && \
    make check && \
    make install && rm -R /tmp/* 

WORKDIR /app

# Install protoc-gen-go
RUN go get github.com/golang/protobuf/protoc-gen-go
RUN cp /go/bin/protoc-gen-go /export/bin/

# Export dependencies
RUN cp /usr/lib/libstdc++* /export/lib/
RUN cp /usr/lib/libgcc_s* /export/lib/

RUN git clone https://github.com/googleapis/googleapis
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get github.com/favadi/protoc-go-inject-tag

RUN cp  /export/bin/* /usr/bin/

 
#copy export folder to /usr in order to use protoc in a fresh Alpine image