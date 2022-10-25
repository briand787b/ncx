# NOTE: Docker build context must be set to the svc/ directory
FROM golang:1.19.2 AS builder
WORKDIR /go/src/gateway-web
COPY gateway-web/go.mod gateway-web/go.sum /go/src/gateway-web/
COPY pkg/go.mod pkg/go.sum /go/src/pkg/
RUN go mod download
COPY . /go/src/
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o /go/bin/app /go/src/gateway-web/cmd

# RUN STAGE
FROM alpine:3.16
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /root
COPY --from=builder /go/bin/app /root/
CMD [ "/root/app" ]