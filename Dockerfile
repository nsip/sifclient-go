###########################
# INSTRUCTIONS
############################
# BUILD: docker build -t nsip/sifclient-go:v0.1.1 .
# TEST: docker run -it -p3000:3000 nsip/sifclient-go:v0.1.1 .
# RUN: docker run -d -p3000:3000 nsip/sifclient-go:v0.1.1
#
###########################
# EXAMPLE DOCUMENTATION
############################
#TODO
# To this example:
#   * Client and Server - both using SSL (server use host certificates)
# To write:
#   * Why multi layers
#   * How the CA certificates - link to the url I find this info
#   * Examples of where this has been used (e.g. sif2json)
#   * Link to the hub.docker official golang docs
#   * Conventions for micro web services.

###########################
# STEP 0 Get them certificates
############################
FROM alpine:latest as certs
RUN apk --no-cache add ca-certificates

############################
# STEP 1 build executable binary (go.mod version)
############################
FROM golang:1.14-stretch as builder
RUN mkdir -p /build
WORKDIR /build
COPY . .
WORKDIR cmd/hitsproxy
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/server

############################
# STEP 2 build a small image
############################
FROM debian:stretch
COPY --from=builder /go/bin/server /go/bin/server
# NOTE - make sure it is the last build that still copies the files
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /go/bin
CMD ["/go/bin/server"]
