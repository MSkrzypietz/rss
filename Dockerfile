FROM golang:1.23 AS backend-builder
WORKDIR /app
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend/ .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o rss

FROM node:22-alpine AS frontend-builder
WORKDIR /app
COPY frontend/package*.json ./
RUN npm ci
COPY frontend/ .
RUN npm run build

FROM alpine:latest
ENV APP_ENV=production
ENV HTTP_PORT=8080
EXPOSE 8080
WORKDIR /app
RUN addgroup -S rss && adduser -S rss -G rss
COPY --chown=rss:rss --from=backend-builder /app/rss .
COPY --chown=rss:rss --from=frontend-builder /app/dist ./static
USER rss
CMD ["/app/rss"]
