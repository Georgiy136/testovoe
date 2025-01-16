package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"myapp/config"
)

func New(cfg config.Postgres) (*bun.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Dbname, cfg.Sslmode)

	connection, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - sql.Open: %w", err)
	}
	if err = connection.Ping(); err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - Ping: %w", err)
	}
	Bun := bun.NewDB(connection, pgdialect.New())

	return Bun, nil
}
