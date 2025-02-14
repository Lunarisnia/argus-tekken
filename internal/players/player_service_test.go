package players

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/Lunarisnia/argus-tekken/database/repo"
	"github.com/Lunarisnia/argus-tekken/internal/db"
	"github.com/Lunarisnia/argus-tekken/internal/players/playermodels"
	"github.com/jackc/pgx/v5"
)

func Test_InsertNewPlayer(t *testing.T) {
	t.Run("Not found", func(t *testing.T) {
		dbURL := "postgres://postgres:password@localhost:5432/argus_db?sslmode=disable"
		conn, err := db.Connect(context.Background(), dbURL)
		if err != nil {
			t.Fail()
		}

		q := repo.New(conn)
		srv := NewPlayerService(q)

		err = srv.InsertNewPlayer(context.Background(), playermodels.Player{
			PolarisID: "12",
		})
		if err != nil {
			fmt.Println("Err: ", err, errors.Is(err, pgx.ErrNoRows))
		}
	})
}
