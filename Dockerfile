# ビルド用のステージ
FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o kakeibo

# 実行用のステージ
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/kakeibo .
EXPOSE 8080
CMD ["./kakeibo"]
