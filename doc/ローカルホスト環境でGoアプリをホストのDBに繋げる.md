## ローカルホスト環境でGoアプリをホストのDBに繋げる

ローカルホスト環境で、Goアプリケーションがコンテナ内で動作している場合、コンテナ内からローカルホスト（127.0.0.1）にアクセスする方法について少し工夫が必要です。コンテナ内からローカルマシンのPostgreSQLにアクセスするためには、以下のいくつかの点を確認する必要があります。

### 1. ホスト名の設定
コンテナ内からローカルのlocalhost（127.0.0.1）に直接接続しようとすると、コンテナ自身のネットワークを指してしまうため、ホストマシンのPostgreSQLにはアクセスできません。

そのため、コンテナからホストマシンのPostgreSQLに接続するには、localhostではなく、ホストマシンのIPアドレスまたは特別なホスト名（例えばhost.docker.internal）を使用する必要があります。

解決策1: host.docker.internal を使用する
host.docker.internalは、Dockerが提供する特別なホスト名で、コンテナ内からホストマシンにアクセスするための名前です。これをdsnのhostに指定することで、コンテナからホストマシンのPostgreSQLに接続できます。

```go
dsn := fmt.Sprintf("user=%s password=%s host=host.docker.internal port=%s dbname=%s sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Port, dbConfig.Name)
```

これで、コンテナ内からホストマシンのPostgreSQLに接続できるはずです。

### 2. PostgreSQLの設定を確認
ローカルのPostgreSQLは、デフォルトでlocalhostまたは127.0.0.1からの接続しか受け入れない設定になっている場合があります。コンテナから接続する場合、pg_hba.confファイルやpostgresql.confを適切に設定する必要があります。

1) pg_hba.confの設定
pg_hba.confファイルで、コンテナのIPアドレス（またはhost.docker.internal）からの接続を許可するように設定します。

```ini
# 例: host.docker.internalからの接続をmd5認証で許可
host    all             all             host.docker.internal/32            md5
```

設定を変更した後、PostgreSQLを再起動します。

```bash
sudo service postgresql restart
```

2) postgresql.confの設定
listen_addressesが適切に設定されているか確認してください。

``` ini
# *を設定することで、任意のIPアドレスから接続を受け入れる
listen_addresses = '*'
```

設定後、再起動して設定を反映します。

```bash
コードをコピーする
sudo service postgresql restart
```

### 3. コンテナのネットワーク設定の確認
コンテナ内からhost.docker.internalにアクセスできない場合、Dockerネットワークの設定に問題があるかもしれません。特に、bridgeネットワークを使用している場合、host.docker.internalへのアクセスがうまくいかないことがあります。この場合、コンテナとホストが同じネットワークにいることを確認してください。

### 4. 接続確認
設定を変更した後、Goのアプリケーション内から接続が成功するか確認してください。

```bash
コードをコピーする
# PostgreSQLが動作しているか確認
psql -h host.docker.internal -U postgres -d demo
```

もし接続できる場合、Goアプリケーションも同様に接続できるはずです。

## まとめ
コンテナ内のGoアプリケーションからローカルのPostgreSQLに接続するためには、host.docker.internalを使ってホストマシンに接続するように設定を変更する必要があります。また、PostgreSQLの設定ファイルで適切な接続設定を行い、ネットワーク接続の確認も忘れずに行ってください。
