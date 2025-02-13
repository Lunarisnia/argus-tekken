package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Lunarisnia/argus-tekken/internal/wank/wankmodels"
)

func main() {
	_, err := getReplays(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// TODO: Set up database
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
