FROM golang:buster as builder 
WORKDIR /tmp
RUN apt-get update && apt-get install -y \
  curl \
  unzip \
  git \
  && rm -rf /var/lib/apt/lists/*
ENV PB_REL="https://github.com/protocolbuffers/protobuf/releases"
RUN curl -LO $PB_REL/download/v3.13.0/protoc-3.13.0-linux-x86_64.zip && unzip protoc-3.13.0-linux-x86_64.zip -d /usr/local
ENV PATH="$PATH:/usr/local/bin"
ENV GO111MODULE=on  
RUN go get google.golang.org/protobuf/cmd/protoc-gen-go \
         google.golang.org/grpc/cmd/protoc-gen-go-grpc
WORKDIR /tmp
RUN git clone https://github.com/googleapis/googleapis  
RUN mkdir /app && cp -r /tmp/googleapis/google /app/google          
# install GRPC-web plugin
WORKDIR /tmp
RUN curl -LO https://github.com/grpc/grpc-web/releases/download/1.2.1/protoc-gen-grpc-web-1.2.1-linux-x86_64         
RUN mv ./protoc-gen-grpc-web-1.2.1-linux-x86_64    /usr/local/bin/protoc-gen-grpc-web && chmod +x /usr/local/bin/protoc-gen-grpc-web

# install ts-protoc-gen
RUN curl -fsSL https://deb.nodesource.com/setup_14.x | bash -
RUN apt-get install -y nodejs
WORKDIR /app
RUN npm install ts-protoc-gen
RUN ls -la ./node_modules
RUN protoc --version

