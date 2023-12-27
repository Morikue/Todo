package main

import (
	"database/sql"
	"fmt"
	"todo/pkg/postgresql"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"

	"todo/config"
	"todo/migrations"
)

func main() {
	// загрузить конфигурацию из переменных среды
	cfg := config.NewMigrationsFromEnv()

	// открыть соединение с базой данных
	db, err := sql.Open("pgx", formDbURI(cfg.Postgres))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// загрузить файлы миграций
	goose.SetBaseFS(migrations.MigrationsFS)
	err = goose.SetDialect("postgres")
	if err != nil {
		panic(err)
	}

	// выполнить из миграций код из раздела -- +goose Up
	err = goose.Up(db, ".", goose.WithAllowMissing())
	if err != nil {
		panic(err)
	}
}

// небольшой хэлпер для формирования DSN для соединения с базой данных
func formDbURI(conf postgresql.PostgreSQL) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&connect_timeout=10",
		conf.PostgresqlUser,
		conf.PostgresqlPassword,
		conf.PostgresqlHost,
		conf.PostgresqlPort,
		conf.PostgresqlDBName,
	)
}
