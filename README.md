# Easy Go Monitor プロジェクト

## このプロジェクトについて

このプロジェクトは外形監視を GUI から気軽にできるようにする Golang 製のプロジェクトです。

## ER 図

![alt text](<easy go monitor.svg>)

## データベースマイグレーション

このプロジェクトは [golang-migrate](https://github.com/golang-migrate/migrate) を使用して
PostgreSQL のマイグレーションを管理しています。

### インストール

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### 新しいマイグレーションファイルの作成

```bash
migrate create -ext sql -dir db/migrations -seq easy-go-monitor-db
```

このコマンドで以下のようなファイルが生成されます:

```
db/migrations/
  0001_create_users_and_monitors.up.sql
  0001_create_users_and_monitors.down.sql
```

### マイグレーションの適用

```bash
migrate -path db/migrations \
  -database "postgres://user:password@localhost:55432/monitor_db?sslmode=disable" up
```

### マイグレーションのロールバック

直近の 1 つを戻す:

```bash
migrate -path db/migrations \
  -database "postgres://user:password@localhost:55432/monitor_db?sslmode=disable" down 1
```

すべて戻す:

```bash
migrate -path db/migrations \
  -database "postgres://user:password@localhost:55432/monitor_db?sslmode=disable" down
```

状態リセット

```bash
migrate -path db/migrations \
  -database "postgres://user:password@localhost:55432/monitor_db?sslmode=disable" drop -f
```

### 確認

PostgreSQL にログインしてテーブルを確認:

```sql
\d users;
\d monitors;
```
