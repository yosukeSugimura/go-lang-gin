# go-lang-gin

`go-lang-gin` は、Go言語とGinフレームワークを使用したプロジェクトです。このプロジェクトは、日本語の姓名判断に基づいて画数を取得し、五格（天格、人格、地格、外格、総格）を計算するAPIを提供します。

## 特徴

- 各文字の画数を取得
- 姓名の五格（天格、人格、地格、外格、総格）の計算
- RESTful API構成
- Ginフレームワーク使用
- Dockerでの簡単なデプロイ対応

## 目次
- [インストール](#インストール)
- [環境変数](#環境変数)
- [使用方法](#使用方法)
- [APIエンドポイント](#APIエンドポイント)
- [Dockerによる実行](#Dockerによる実行)
  
## インストール

### 1. リポジトリをクローン

```bash
git clone https://github.com/yosukeSugimura/go-lang-gin.git
cd go-lang-gin
```
### 1. 必要なパッケージをインストール

```bash
go mod tidy
```

### 1. .envファイルの作成

- ルートディレクトリに .env ファイルを作成し、環境変数を設定します。

```plaintext
PORT=8080
GIN_MODE=debug
DB_HOST=localhost
DB_PORT=5432
DB_USER=username
DB_PASSWORD=password
DB_NAME=database_name
```

## 環境変数

| 変数名 | 説明 | デフォルト値 |
| ---- | ---- |---- |
| `PORT`	      | サーバーのポート番号	          |8080|
| `GIN_MODE`	  | Ginの実行モード（debug/release） |debug|
| `DB_HOST`	  | データベースのホスト	          |localhost|
| `DB_PORT`	  | データベースのポート番号	       |5432|
| `DB_USER`	  | データベースのユーザー名	       |-|
| `DB_PASSWORD` | データベースのパスワード	       |-|
| `DB_NAME`	  | データベース名	                 |-|

## 使用方法

### サーバーの起動

- 以下のコマンドでサーバーを起動できます。

```bash
go run cmd/main.go
```

- サーバーが起動すると、デフォルトでは http://localhost:8080 でアクセスできます。

## APIエンドポイント
### 1. ルートページ
- URL: /
- メソッド: GET
- 説明: ホームページ（index.html）を表示
### 2. 各文字の画数を取得
- URL: /seimei/:name/:sei
- メソッド: GET
- パラメータ:
  - name: 名前
  - sei: 姓
- レスポンス:
```json
{
    "stroke_counts": [
        { "character": "姓", "strokes": 10 },
        { "character": "名", "strokes": 12 }
    ]
}
```
## 3. 五格を計算
- URL: /seimei/:name/:sei/grids
- メソッド: GET
- パラメータ:
  - name: 名前
  - sei: 姓
- レスポンス:
```json
{
    "tenkaku": 10,
    "jinkaku": 12,
    "chikaku": 8,
    "gaikaku": 7,
    "sokaku": 37
}
```
## Dockerによる実行
### 1. Dockerイメージのビルド

```bash
docker build -t go-lang-gin .
```

### 2. Dockerコンテナの実行
```bash
docker run -p 8080:8080 go-lang-gin
```
- Dockerで実行すると、ホストマシンの http://localhost:8080 でAPIを利用できます。

---

## ライセンス

このプロジェクトは MIT ライセンスのもとで公開されています。

## 作者
Yosuke Sugimura


---