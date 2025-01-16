package app

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
)

func migrateUp(conn *bun.DB) {
	// Накатываем миграции
	driver, err := database.Open(conn.String())
	if err != nil {
		log.Fatal().Err(err).Msg("не удалось инициализировать систему миграций Postgres")
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
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
