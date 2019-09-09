############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/creart_new
COPY . .

# Fetch dependencies.
# Using go get.
RUN go get -d -v
# Build the binary.
RUN go build -o creart_new

############################
# STEP 2 build a small image
############################
FROM scratch

# Run the hello binary.
ENTRYPOINT ["main.go"]