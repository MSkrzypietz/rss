FROM golang:1.22

ENV APP_ENV=production
ENV PORT=8080

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux go build -o /rss

EXPOSE 8080
CMD ["/rss"]
