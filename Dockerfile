FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o lown main.go

FROM alpine:latest
RUN apk add --no-cache git ca-certificates curl
COPY --from=builder /app/lown /usr/local/bin/lown
ENTRYPOINT ["/usr/local/bin/lown"]
