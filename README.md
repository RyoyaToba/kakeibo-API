## 説明

This repository was created for learning the Go programming language. It implements a simple API using Gin.

I will implement using tools like Docker and Docker Compose, as well as the layered architecture of Go.

## 環境構築

#### go install
```
# go install
brew go install
```
#### docker install
```
# docker install
brew install --cask docker
```

#### application start

```
# docker build
docker-compose build --no-cache

# docker up
docker-compose up
```

## ディレクトリ構成

```
.
├── Dockerfile
├── README.md
├── api
│   ├── application
│   ├── interface
│   └── repository
├── doc
│   └── 初期化スクリプトについて.md
├── docker-compose.yaml
├── go.mod
├── go.sum
├── main.go
└── mysql
    ├── Dockerfile
    └── init.sql
```

## API
```
/v1/getmessage/{id}
```
