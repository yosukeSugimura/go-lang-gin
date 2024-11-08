# Go 1.23 の標準イメージを使用
FROM golang:1.23

# 必要なツールのインストール
RUN apt-get update && apt-get install -y git

# Airをソースコードからクローンしてビルド
RUN git clone https://github.com/air-verse/air.git /tmp/air \
    && cd /tmp/air \
    && go build -o /bin/air .

# 作業ディレクトリの設定
WORKDIR /app

# 必要なファイルのコピー
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# プロジェクトファイルのコピー
COPY . .

# Airでアプリケーションを
