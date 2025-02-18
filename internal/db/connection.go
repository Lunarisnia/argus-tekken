package db

import (
	"context"
	"log"
	"os"

	"github.com/Lunarisnia/argus-tekken/database/repo"
	"github.com/jackc/pgx/v5"
)

func Connect(ctx context.Context, url string) (*pgx.Conn, error) {
	// postgres://jack:secret@pg.example.com:5432/mydb
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func ConnectDatabase(ctx context.Context) (*repo.Queries, *pgx.Conn) {
	dbURL := os.Getenv("ARGUS_DB")
	if dbURL == "" {
		log.Fatal("please add db url to ARGUS_DB")
	}

	conn, err := Connect(ctx, dbURL)
	if err != nil {
		log.Fatal(err)
	}

	q := repo.New(conn)

	return q, conn
}
