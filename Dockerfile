###########################
# INSTRUCTIONS
############################
# BUILD: docker build -t nsip/sifclient-go:v0.1.1 .
# TEST: docker run -it -p3000:8089 nsip/sifclient-go:v0.1.1 .
# RUN: docker run -d -p3000:8089 nsip/sifclient-go:v0.1.1
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
# (note, step 2 is using alpine now) 
# FROM alpine:latest as certs

############################
# STEP 1 build executable binary (go.mod version)
############################
FROM golang:1.15.0-alpine3.12 as builder
RUN apk --no-cache add ca-certificates
RUN apk update && apk add git
RUN apk add gcc g++
RUN mkdir -p /build
WORKDIR /build
COPY . .
WORKDIR cmd/hitsproxy
RUN go build -o /build/app

############################
# STEP 2 build a small image
############################
#FROM debian:stretch
FROM alpine
COPY --from=builder /build/app /app
# NOTE - make sure it is the last build that still copies the files
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["./app"]
