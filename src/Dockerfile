FROM golang:1.20-alpine as builder

ARG PROJECT_NAME="app"
ARG VERSION="dev"
ARG COMMIT="none"
ENV CGO_ENABLED=0
ENV GOPROXY=https://proxy.golang.org

WORKDIR /app

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz && mv migrate /usr/local/bin
RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/app_bin .

FROM debian:buster

WORKDIR /app

COPY --from=builder /app/app_bin /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

ENTRYPOINT ["/app/app_bin"]