FROM golang:1.24 AS backend-builder
WORKDIR /app
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend/ .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o rss-api ./cmd/api

FROM alpine:latest
ENV APP_ENV=production
ENV HTTP_PORT=8080
EXPOSE 8080
WORKDIR /app
RUN addgroup -S rss-api && adduser -S rss-api -G rss-api
RUN chown -R rss-api:rss-api /app
COPY --chown=rss:rss --from=backend-builder /app/rss-api .
USER rss-api
CMD ["/app/rss-api", "--logPath", "/app/rss-api.log"]
