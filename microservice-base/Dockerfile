# Basic Alpine image for microservice deployment
FROM alpine:3.13
#RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN apk update && \
    apk add --no-cache tzdata ca-certificates alpine-conf
RUN /sbin/setup-timezone -z Europe/Amsterdam