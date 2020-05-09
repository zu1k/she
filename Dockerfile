FROM golang:alpine as builder

RUN apk add --no-cache git build-base
WORKDIR /she-src
COPY . /she-src
RUN git checkout dist && \
    cp -rf /she-src /dist && \
    git checkout master && \
    go mod download && \
    make linux-amd64 && \
    mv ./bin/she-linux-amd64 /she

FROM alpine:latest

COPY --from=builder /she /
COPY --from=builder /she-src/source/bleveindex/dict/dictionary.txt /
COPY --from=builder /dist /
ENTRYPOINT ["/she"]