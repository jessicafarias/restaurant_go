FROM golang:1.16.3-buster as builder

RUN mkdir /build
WORKDIR /build

ADD . .

RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -gcflags=-trimpath=$PWD -o main .

FROM alpine:3.13

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

ENV GOROOT /usr/local/go
ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip

RUN mkdir /app

COPY --from=builder /build/main /app/

EXPOSE 8080
WORKDIR /app
CMD ["./main"]