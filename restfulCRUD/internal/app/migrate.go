package app

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	postgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
)

func migrateUp(conn *bun.DB) {

	instance, err := postgres.WithInstance(conn.DB, &postgres.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("dStub.WithInstance ошибка при применении миграций Postgres")
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		instance,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("не удалось инициализировать систему миграций Postgres")
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal().Err(err).Msg("ошибка при применении миграций Postgres")
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Info().Msg("нет новых миграций для применения")
		return
	}

	log.Info().Msg("миграции успешно применены")
	// END Накатываем миграции
}
