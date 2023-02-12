# Build Stage
FROM golang:1.19.5-alpine3.17 AS builder

WORKDIR /app

COPY . .

RUN go build -o main cmd/main.go
RUN apk add curl
RUN curl -L -s https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine
WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/migrate /usr/bin/migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY internal/migration ./migration

EXPOSE 8000
CMD ["/app/main"]
ENTRYPOINT ["/app/start.sh"]



# sudo mv migrate /usr/bin/migrate