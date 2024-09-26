# Golangの公式イメージをベースに使用
FROM golang:1.21

# ワーキングディレクトリを作成
WORKDIR /app

# ソースコードをコンテナにコピー
COPY . .

# アプリケーションをビルド
RUN go build -o /hello-world-app

# ポート8080を開放
EXPOSE 8080

# アプリケーションを起動
CMD ["/hello-world-app"]
