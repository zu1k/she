FROM golang:alpine as builder

RUN apk add --no-cache make git
WORKDIR /she-src
COPY . /she-src
RUN go mod download && \
    make linux-amd64 && \
    mv ./bin/she-linux-amd64 /she

FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /she /
ENTRYPOINT ["/she"]