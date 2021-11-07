FROM golang:1.17-alpine3.14 AS build-env

# Set environment variables
ENV GO_WORKDIR /go/src/github.com/i-sevostyanov/NanoDB
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

# Put everything together in a clean image
FROM alpine:3.14

# Copy binary into PATH
COPY --from=build-env /go/bin/repl /usr/local/bin/repl

ENTRYPOINT ["repl"]
