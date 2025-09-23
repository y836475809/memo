# ベースイメージを指定
FROM golang:1.24.6 AS builder

# ワーキングディレクトリを設定
# WORKDIR /usr/src/app
WORKDIR /app

# COPY go.mod go.* ./
# RUN go mod download && go mod verify
RUN go install github.com/go-delve/delve/cmd/dlv@latest
# RUN apt-get update && apt-get install -y --no-install-recommends gdb

# アプリケーションをビルド
COPY . .
RUN --mount=type=bind,target=. go build -gcflags='all=-N -l' -ldflags=-compressdwarf=false -o /bin/go-test1
# COPY . .
# RUN go build -gcflags="all=-N -l" -v -o /usr/local/bin/app ./...

# ワーキングディレクトリを設定
WORKDIR /bin
# RUN echo "add-auto-load-safe-path /usr/local/go/src/runtime/runtime-gdb.py" > /root/.gdbinit

# アプリケーションを実行
CMD ["./go-test1"]
# CMD ["./app"]
# CMD ["/go/bin/dlv", "exec" "./go-test1"]
# CMD ["dlv","debug"]

# ポート8080を公開
EXPOSE 8080