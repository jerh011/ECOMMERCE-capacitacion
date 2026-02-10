# Comandos necesarios para correr el proyecto
  - go mod init EcomerceProject
  - go get -u github.com/go-chi/chi/v5
  - go mod tidy
  - go run ./cmd
  - go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
  - sqlc generate
  - go install github.com/pressly/goose/v3/cmd/goose@latest
  - go get github.com/jackc/pgx/v5

  goose -s create add_some_column sql

  minuto 1:24:31 
  problemas al conectarte con la base de datos  (rewsolver)
