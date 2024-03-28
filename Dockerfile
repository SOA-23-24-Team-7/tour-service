FROM golang:alpine AS builder
WORKDIR /app
COPY ./RestServiceGo/WebServerWithDB .
EXPOSE 8087
ENTRYPOINT ["go", "run", "main.go"]