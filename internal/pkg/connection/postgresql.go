package main

import (
	"context"
	"fmt"
	"os"
	"net/url"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgreSQL(conf *Configuration)(*pgxpool.Pool, error) {
	get := func(v string) string {
		res, err := conf.Get(v)
		if err != nil {
			log.Fatalf("Couldn't get configuration value %s: %s", v, err)
		}

		return res
	}

	// FIXME: with an elegant solution
	databaseHost := get("DATABASE_HOST")
	databasePort := get("DATABASE_PORT")
	databaseUsername := get("DATABASE_USERNAME")
	databasePassword := get("DATABASE_PASSWORD")
	databaseName := get("DATABASE_NAME")
	databaseSSLMode := get("DATABASE_SSLMODE")

	dsn := url.URL {
		Scheme: "postgres",
		User: url.UserPassword(databaseUsername, databasePassword)
		Host: fmt.Sprintf("%s:%s", databaseHost, databaseHost)
		Path: databaseName
	

	q := dsn.Query()
	q.Add("sslMode", databaseSSLMode)

	dsn.RawQuery = q.Encode()

	dbpool, err := pgxpool.Connect(context.Background(), dsn.String())
	if err != nil {
		return nil, internal.WrapError(err, internal.ErrorCodeUnknown, "pgxPool.Connect")
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, internal.WrapError(err, internal.ErrorCodeUnknown, "db.Ping")
	}
	
	return pool, nil
}
