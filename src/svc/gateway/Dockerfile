# NOTE: Docker build context must be set to the svc/ directory
FROM golang:1.18 AS builder
WORKDIR /go/src/gateway
COPY gateway/go.mod gateway/go.sum /go/src/gateway/
COPY pkg/go.mod pkg/go.sum /go/src/pkg/
RUN go mod download
COPY . /go/src/
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o /go/bin/app /go/src/gateway/cmd

# RUN STAGE
FROM alpine:3.16
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /root
COPY --from=builder /go/bin/app /root/
CMD [ "/root/app" ]