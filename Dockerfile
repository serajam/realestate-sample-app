# syntax = docker/dockerfile:1
FROM golang:1.21-alpine as builder

RUN apk add --no-cache ca-certificates

ARG BUILD_HASH=dev
ARG BINARY_NAME=realestate
ENV DIST=/go/src/realestate

RUN mkdir -p ${DIST}
WORKDIR ${DIST}

# enable Go modules support
ENV GO111MODULE=on

COPY cmd ./cmd
COPY internal ./internal
COPY docs ./docs
COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ./dist/${BINARY_NAME} -x \
      -ldflags="-X main.build=${BUILD_HASH} -s -w -extldflags '-static'" ./cmd/estate

FROM alpine:3.19

# Set environment variables
ENV USER=estate
ENV UID=1000
ENV GID=1000
ENV HOME=/home/$USER

# Add new user and group
RUN addgroup -g $GID $USER && \
    adduser -u $UID -G $USER -h $HOME -D $USER && \
    chown $USER:$USER $HOME

# Install necessary packages
RUN apk update && \
    apk add --no-cache \
        ca-certificates \
    rm -rf /var/cache/apk/*

ARG BINARY_NAME=realestate
ENV DIST=/go/src/realestate
ENV BINARY="./${BINARY_NAME}"

WORKDIR $HOME

COPY --from=builder ${DIST}/dist/${BINARY_NAME} ./
RUN chown $USER:$USER $BINARY_NAME

# Switch to new user
USER $USER

ENTRYPOINT $BINARY