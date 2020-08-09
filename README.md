# Job Survey

## 実行環境

- macOS 10.14.6 以上
- Docker Desktop 2.3.0.4 (Engine 19.03.12)
- Golang
- MySQL
- 楽天ブックス書籍検索 API

## 環境構築

- build/env/ に survey.env と mysql.env を配置する
- docker-compose でコンテナを立ち上げる

```zsh
% docker-compose up --build -d
```

## 求人情報取得

- API から書籍情報を取得する
- assets/ に tsv ファイルを出力する

```zsh
% docker-compose exec survey cmd/survey
```
