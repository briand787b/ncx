FROM golang:1.18 AS builder
WORKDIR /go/src
COPY go.mod go.sum /go/src/
RUN go mod download
COPY . /go/src/
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o /go/bin/app /go/src/cmd

# RUN STAGE
FROM alpine:3.16
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /root
COPY --from=builder /go/bin/app /root/
CMD [ "/root/app" ]