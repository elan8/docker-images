FROM golang:1.14-alpine as builder 
WORKDIR /tmp
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
WORKDIR /tmp     
RUN git clone https://github.com/googleapis/googleapis   
#RUN ls /export/lib

FROM golang:1.14-alpine 
WORKDIR /golang
COPY tools.go .
COPY go.mod .
COPY go.sum .
RUN go install \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
    github.com/golang/protobuf/protoc-gen-go

#COPY --from=builder /export /export
#RUN cp  /export/bin/* /usr/bin/
COPY --from=builder /export/lib /usr/lib
COPY --from=builder /export/bin /usr/bin
COPY --from=builder /export/include/google /usr/include/google
COPY --from=builder /tmp/googleapis/google /app/google
RUN ls /usr/lib
WORKDIR /app
RUN  apk add libstdc++ clang

#RUN cp  /export/bin/* /usr/bin/
#RUN cp -r /export/include/google /usr/include/
