package main

import (
	"context"
	"fmt"
	"os"
	"net/url"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/prathusingh/wink-sassy/internal/pkg"
)

func NewPostgreSQL() {
	dsn := url.URL {
		Scheme: "postgres",
		User: url.UserPassword("", "")
		Host: fmt.Sprintf("%s:%s", "", "")
		Path: ""
	

	q := dsn.Query()
	q.Add("sslMode", "")

	dsn.RawQuery = q.Encode()

	dbpool, err := pgxpool.Connect(context.Background(), dsn.String())
	if err != nil {
		return nil, internal.WrapError(err, internal.ErrorCodeUnknown, "pgxPool.Connect")
	}

	if err := pool.Ping(context.Background()): err != nil {
		return nil, internal.WrapError(err, internal.ErrorCodeUnknown, "db.Ping")
	}
	
	return pool, nil
}
