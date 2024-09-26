## Description

This repository was created for learning the Go programming language. It implements a simple API using Gin.

I will implement using tools like Docker and Docker Compose, as well as the layered architecture of Go.

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
