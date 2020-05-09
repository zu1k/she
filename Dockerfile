FROM golang:alpine as builder

RUN apk add --no-cache git build-base
WORKDIR /she-src
COPY . /she-src
RUN go mod download && \
    make linux-amd64 && \
    mv ./bin/she-linux-amd64 /she

FROM alpine:latest

COPY --from=builder /she /
COPY --from=builder /she-src/source/bleveindex/dict/dictionary.txt /
ENTRYPOINT ["/she"]