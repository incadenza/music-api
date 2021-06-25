## Setup
1) Ensure Golang is installed: https://golang.org/dl/

1) Download provided dataset into a postgres database:  https://gist.github.com/shahidhk/351f7201c9cc35be5fd9f40e48113637/raw/0692054166afb79c2c681b680e6c52dbdf65f95a/chinook_postgres.sql

2) `go mod download` in the project directory

2) Run in development mode with `go run . -dsn=$LOCAL_PG_DSN -trusted-origins=http://localhost:3000` where `$LOCAL_PG_DSN` is the connection string to a running PostgreSQL database. By default, the server will be listening on port `8080`.

## Todos
- Structured logging: info, warning, error, etc
- Integration testing setup, using something like https://github.com/ory/dockertest
- Makefile to automate build and testing tasks
- Perhaps some simple sorting, filtering & searching for tracks (there's quite a few).