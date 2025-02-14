package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/Lunarisnia/argus-tekken/internal/db"
	"github.com/Lunarisnia/argus-tekken/internal/wank/wankmodels"
	"github.com/jackc/pgx/v5"
)

var (
	dbURL string
	conn  *pgx.Conn
)

func init() {
	dbURL = os.Getenv("ARGUS_DB")
	if dbURL == "" {
		log.Fatal("please add your database url to env (ARGUS_DB)")
	}

	c, err := db.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatal(err)
	}

	conn = c
}

func main() {
	_, err := getReplays(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

func getReplays(ctx context.Context) ([]wankmodels.Replay, error) {
	client := http.Client{}

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"https://wank.wavu.wiki/api/replays",
		nil,
	)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		err := errors.New("non 200 response")
		log.Println("Error Response: ", resp.Body)
		return nil, err
	}

	replays := make([]wankmodels.Replay, 0)

	err = json.NewDecoder(resp.Body).Decode(&replays)
	if err != nil {
		return nil, err
	}

	return replays, nil
}
