FROM golang:1.20.1-alpine3.17 AS builder

LABEL maintainer="Muhammad Abdurrahman <rachman.sd@gmail.com> (https://github.com/rachmanzz)"

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o botservice .

FROM scratch

# Copy binary and config files from /build to root folder of scratch container.
COPY --from=builder ["/build/botservice",  "/"]

# Command to run when starting the container.
ENTRYPOINT ["/botservice"]