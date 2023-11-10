package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"hvalfangst/golang-fasthttp-pgx/src/configuration"
)

func CreateDB(config configuration.Db) (*pgxpool.Pool, error) {
	connConfig, err := pgxpool.ParseConfig(config.DataSourceName)
	if err != nil {
		return nil, err
	}

	connConfig.MaxConns = config.MaxConns
	connConfig.MinConns = config.MinConns

	db, err := pgxpool.ConnectConfig(context.Background(), connConfig)
	if err != nil {
		return nil, err
	}

	return db, nil
}
