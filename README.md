# Easy Go Monitorプロジェクト

## このプロジェクトについて

このプロジェクトは

## ER図

![alt text](<er_easy-go-monitor.svg>)

## Database Migration

このプロジェクトは [golang-migrate](https://github.com/golang-migrate/migrate) を使用して
PostgreSQL のマイグレーションを管理しています。

### インストール

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### 新しいマイグレーションファイルの作成

```bash
migrate create -ext sql -dir db/migrations -seq create_users_and_monitors
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
  -database "postgres://<user>:<password>@localhost:55432/monitor_db?sslmode=disable" up
```

### マイグレーションのロールバック

直近の1つを戻す:

```bash
migrate -path db/migrations \
  -database "postgres://<user>:<password>@localhost:55432/monitor_db?sslmode=disable" down 1
```

すべて戻す:

```bash
migrate -path db/migrations \
  -database "postgres://<user>:<password>@localhost:55432/monitor_db?sslmode=disable" down
```

### 確認

PostgreSQLにログインしてテーブルを確認:

```sql
\d users;
\d monitors;
```
