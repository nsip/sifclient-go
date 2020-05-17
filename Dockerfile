#
# https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324
#
# BUILD: docker build -t nsip/hitsproxy .
# TEST: docker run -it -p8089:8089 nsip/hitsproxy
# RUN: docker run -d -p8089:8089 nsip/hitsproxy
############################
# STEP 1 build executable binary
############################
FROM golang:1.13-stretch as builder
COPY . .
RUN go get
RUN go get github.com/labstack/echo/middleware
RUN go get github.com/nsip/sifclient-go
RUN go get github.com/nsip/sifdata-go
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/hitsproxy cmd/hitsproxy/main.go
############################
# STEP 2 build a small image
############################
FROM debian:stretch
COPY --from=builder /go/bin/hitsproxy /go/bin/hitsproxy
CMD ["/go/bin/hitsproxy"]
