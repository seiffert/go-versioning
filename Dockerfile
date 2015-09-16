FROM alpine:3.2

RUN apk --update add ca-certificates

COPY go-versioning /app

ENTRYPOINT ["/app"]
CMD ["--"]
