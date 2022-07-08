FROM golang:1.18-alpine3.15

WORKDIR /build/
COPY . .
RUN apk --no-cache add build-base curl && \
    GOOS=linux go build -o app .

USER 1000
EXPOSE 8082
ENTRYPOINT ["./app"]
