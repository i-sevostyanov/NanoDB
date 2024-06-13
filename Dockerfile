FROM golang:1.22-alpine3.19 AS build-env

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
RUN cd $GO_WORKDIR/cmd/shell && go install

# Put everything together in a clean image
FROM scratch

# Copy binary into PATH
COPY --from=build-env /go/bin/shell /usr/local/bin/shell

CMD ["shell"]
