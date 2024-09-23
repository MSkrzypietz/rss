FROM golang:1.22

ENV APP_ENV=production
ENV HTTP_PORT=8080

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /rss

EXPOSE 8080
EXPOSE 8081
CMD ["/rss"]
