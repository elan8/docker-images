FROM golang:1.16.0-alpine3.12 as builder

RUN apk update && \
    apk add --no-cache tzdata ca-certificates alpine-conf libstdc++ clang git openssh-client && rm -rf /var/cache/apk/*
RUN /sbin/setup-timezone -z Europe/Amsterdam
WORKDIR /tmp
RUN wget https://github.com/bufbuild/buf/archive/v0.39.1.tar.gz 
RUN tar -xzf v0.39.1.tar.gz && rm v0.39.1.tar.gz
WORKDIR /tmp/buf-0.39.1
RUN go mod download
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -trimpath -o /go/bin/buf ./cmd/buf
ENV PATH="/go/bin:${PATH}"

WORKDIR /git

ENV GOOGLEAPIS_COMMIT="d4aa417ed2bba89c2d216900282bddfdafef6128"
RUN git clone https://github.com/googleapis/googleapis && cd googleapis && git reset --hard "${GOOGLEAPIS_COMMIT}"

WORKDIR /workspace

# FROM alpine:3.12.4

# RUN apk add --update --no-cache \
#     ca-certificates \
#     git \
#     openssh-client && \
#   rm -rf /var/cache/apk/*

# COPY --from=builder /go/bin/buf /usr/local/bin/buf

# ENTRYPOINT ["/usr/local/bin/buf"]