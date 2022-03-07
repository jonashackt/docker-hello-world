# We need a golang build environment first
FROM golang:1.17.3-alpine3.13

WORKDIR /go/src/app
ADD hello-world.go /go/src/app

RUN go build hello-world.go

# We use a Docker multi-stage build here in order that we only take the compiled go executable
FROM alpine:3.14

LABEL org.opencontainers.image.source="https://github.com/jonashackt/docker-hello-world"

COPY --from=0 "/go/src/app/hello-world" hello-world

ENTRYPOINT ./hello-world


