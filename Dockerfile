# ベースイメージを指定
FROM golang:1.24.6 AS builder

# ワーキングディレクトリを設定
WORKDIR /app

# アプリケーションをビルド
RUN --mount=type=bind,target=. go build -o /bin/go-test1

# ワーキングディレクトリを設定
WORKDIR /bin

# アプリケーションを実行
CMD ["./go-test1"]

# ポート8080を公開
EXPOSE 8080