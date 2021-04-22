FROM golang:1.16.3-alpine3.13 AS build-env

# Set environment variables
ENV GO_WORKDIR /go/src/github.com/i-sevostyanov/NanoDB
ENV GO111MODULE=on
ENV CGO_ENABLED=0

# Set working directory
WORKDIR $GO_WORKDIR
ADD . $GO_WORKDIR

# Add git and openssh
RUN set -eux; apk update; apk add --no-cache git openssh

# Install dependencies
RUN go mod download
RUN go mod verify
RUN cd $GO_WORKDIR/cmd/repl && go install

# Build ca-certificates
FROM alpine:latest as certs

# Add ca-certificates dependency
RUN apk --update add ca-certificates

# Put everything together in a clean image
FROM alpine:3.13

# Add ca-certificates
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Copy binary into PATH
COPY --from=build-env /go/bin/repl /usr/local/bin/repl

ENTRYPOINT ["repl"]
