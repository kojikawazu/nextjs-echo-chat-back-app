# Stage 1: Build
FROM golang:1.23-alpine AS builder

# 作業ディレクトリを設定
WORKDIR /app

# 必要なパッケージをインストール
RUN apk add --no-cache git

# Go モジュールの依存関係をキャッシュしておく
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# ビルド
RUN go build -o main main.go

# Stage 2: Run
FROM alpine:latest

# 作業ディレクトリを設定
WORKDIR /app

# 依存関係をインストール
RUN apk add --no-cache ca-certificates

# ビルド済みのバイナリをコピー
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# 環境変数を使ってポートを動的に変更
EXPOSE ${NGINX_APP_PORT}

# サーバーを実行
CMD ["./main"]
