package main

import (
	"context"
	"fmt"
	"os"
	"net/url"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgreSQL() {
	dsn := url.URL {
		Scheme: "postgres",
		User: url.UserPassword("", "")
		Host: fmt.Sprintf("%s:%s", "", "")
		Path: ""
	}

	q := dsn.Query()
	q.Add("sslMode", "")

	dsn.RawQuery = q.Encode()
	
	dbpool, err := pgxpool.Connect(context.Background(), dsn.String())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()
}
