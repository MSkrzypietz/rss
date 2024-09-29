FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o rss

FROM alpine:latest
ENV APP_ENV=production
ENV HTTP_PORT=8080
EXPOSE 8080
WORKDIR /app
RUN addgroup -S rss && adduser -S rss -G rss
COPY --chown=rss:rss --from=builder /app/rss .
USER rss
CMD ["/app/rss"]
