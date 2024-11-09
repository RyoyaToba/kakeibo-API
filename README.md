## Description

このレポジトリでは[Golang](https://go.dev/)を利用したAPIの練習を行なっています. 

フレームワークは[Gin Framework](https://gin-gonic.com/ja/docs/)を利用しています.

## Tool

comming soon

## Environment Setup

#### go install
```
# go install
brew go install

# install check
go version
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

## Directory Structure

```
.
├── README.md
├── api
│   ├── Dockerfile
│   ├── application
│   ├── interface
│   ├── openAPI.yml
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

## API Specification

The API definition is written using OpenAPI. Please check the specifications using Swagger Editor.

[openAPI.yml](./api/openAPI.yml)  
[Swagger Editor](https://editor.swagger.io/)

## What is Layered Architecture?

Please refer to the following for details on the layered architecture.

[今すぐ「レイヤードアーキテクチャ+DDD」を理解しよう。（golang）](https://qiita.com/tono-maron/items/345c433b86f74d314c8d)
