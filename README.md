# Job Survey

## 実行環境

- macOS 10.14.6 以上
- Docker Desktop 2.3.0.4 (Engine 19.03.12)
- Golang
- MySQL

## 環境構築

- build/env/ に survey.env と mysql.env を配置する
- docker-compose でコンテナを立ち上げる

```zsh
% docker-compose up --build -d
```
