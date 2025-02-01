FROM golang:1.23.5 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gorestcontroller .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/gorestcontroller .

EXPOSE 8080

RUN chmod +x gorestcontroller

CMD ["./gorestcontroller"]