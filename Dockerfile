FROM golang:1.22

ARG GOOS=linux
ARG GOARCH=amd64

ENV APP_ENV=production
ENV HTTP_PORT=8080
ENV HTTPS_PORT=8081

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-s -w" -o /rss

EXPOSE 8080
EXPOSE 8081
CMD ["/rss"]
