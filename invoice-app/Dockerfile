FROM golang:1.18-alpine3.15

WORKDIR /build/
COPY . .
RUN apk add build-base && \
    GOOS=linux && \
    go build -o app .

USER 1000
ENTRYPOINT ["./app"]
