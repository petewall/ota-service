FROM golang:1.19-alpine AS builder

WORKDIR /ota-service

COPY . /ota-service/

ARG GOOS=linux
ARG GOARCH=amd64

RUN apk add bash make
RUN make build

FROM alpine

WORKDIR /

COPY --from=builder /ota-service/build/ota-service /ota-service

ENTRYPOINT ["/ota-service"]
