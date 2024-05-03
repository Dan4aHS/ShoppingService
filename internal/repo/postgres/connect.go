package postgres

import (
	"ShoppingExpensesService/internal/config"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

const driverName = "pgx"

func ConnectPostgres(cfg *config.Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Database,
	)
	connCfg, err := pgx.ParseConfig(dsn)
	if err != nil {
		slog.Error("failed to parse postgres connection string", "error", err)
		return nil, err
	}
	db, err := sql.Open(driverName, stdlib.RegisterConnConfig(connCfg))
	if err != nil {
		slog.Error("failed to open postgres connection", "error", err)
		return nil, err
	}
	dbx := sqlx.NewDb(db, driverName)
	err = db.Ping()
	if err != nil {
		slog.Error("failed to ping postgres connection", "error", err)
		return nil, err
	}
	slog.Info("successfully connected to postgres database")
	return dbx, nil
}
