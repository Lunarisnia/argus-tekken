package db

import (
	"context"

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
