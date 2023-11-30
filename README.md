# vehicle-signal-decoding

Api for managing vehicle signal decoding on the DIMO platform.
## Developing locally

To download postgres:

```bash
brew install postgresql
pg_ctl -D /usr/local/var/postgres start && brew services start postgresql
```
If postgres is already installed, go ahead and create database in postgres. 

```bash
psql postgres
create user dimo PASSWORD 'dimo';
grant dimo to postgres;
create database vehicle_signal_decoding_api
    with owner dimo;
```

Open postgres database in DataGrip to view schema, tables, etc.

**TL;DR**

```bash
cp settings.sample.yaml settings.yaml
docker compose up -d
go run ./cmd/vehicle-signal-decoding migrate
go run ./cmd/vehicle-signal-decoding
```

### Regenerate swagger docs

`swag init -g cmd/vehicle-signal-decoding/main.go --parseDependency --parseInternal --generatedTime true`

## Generating gRPC client and server code

1. Install the protocol compiler plugins for Go using the following commands

```bash
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

2. Run protoc in the root directory

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pkg/grpc/*.proto
```

## Linting

`brew install golangci-lint`

`golangci-lint run`

This should use the settings from `.golangci.yml`, which you can override.

If brew version does not work, download from https://github.com/golangci/golangci-lint/releases (darwin arm64 if M1), then copy to /usr/local/bin and sudo xattr -c golangci-lint

### Database ORM

This is using [sqlboiler](https://github.com/volatiletech/sqlboiler). The ORM models are code generated. If the db changes,
you must update the models.

Make sure you have sqlboiler installed:

```bash
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
```

To generate the models:

```bash
sqlboiler psql --no-tests --wipe
```

_Make sure you're running the docker image (ie. docker compose up)_

If you get a command not found error with sqlboiler, make sure your go install is correct.
[Instructions here](https://jimkang.medium.com/install-go-on-mac-with-homebrew-5fa421fc55f5)

### Adding migrations

To install goose in GO:
```bash
$ go get github.com/pressly/goose/v3/cmd/goose@v3.5.3
export GOOSE_DRIVER=postgres
```

To install goose CLI:
```bash
$ go install github.com/pressly/goose/v3/cmd/goose
export GOOSE_DRIVER=postgres
```

Have goose installed, then:

`goose -dir internal/infrastructure/db/migrations create slugs-not-null sql`

## Local development

Importing data: Device definition exports are [here]([url](https://drive.google.com/drive/u/1/folders/1WymEqZo-bCH2Zw-m5L9u_ynMSwPeEARL))
You can use sqlboiler to import or this command:
```sh
psql "host=localhost port=5432 dbname=vehicle_signal_decoding_api user=dimo password=dimo" -c "\COPY vehicle_signal_decoding_api.integrations (id, type, style, vendor, created_at, updated_at, refresh_limit_secs, metadata) FROM '/Users/aenglish/Downloads/drive-download-20221020T172636Z-001/integrations.csv' DELIMITER ',' CSV HEADER"
```

### Starting Kafka locally

`$ brew services start kafka`
`$ brew services start zookeeper`

This will use the brew services to start kafka locally on port 9092. One nice thing of this vs. docker-compose is that we can use this 
same instance for all our different locally running services that require kafka. 

### Produce some test messages

`$ go run ./cmd/test-producer`

In current state this only produces a single message, but should be good enough starting point to test locally. 

### Create decoding topic 

`kafka-topics --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic topic.JOB.decoding`

### Sample read messages in the topic

`kafka-console-consumer --bootstrap-server localhost:9092 --topic topic.JOB.decoding --from-beginning`