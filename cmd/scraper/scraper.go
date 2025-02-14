package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/Lunarisnia/argus-tekken/database/repo"
	"github.com/Lunarisnia/argus-tekken/internal/db"
	"github.com/Lunarisnia/argus-tekken/internal/players/playermodels"
	"github.com/Lunarisnia/argus-tekken/internal/wank/wankmodels"
)

var (
	dbURL   string
	queries *repo.Queries
)

func init() {
	dbURL = os.Getenv("ARGUS_DB")
	if dbURL == "" {
		log.Fatal("please add your database url to env (ARGUS_DB)")
	}

	conn, err := db.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatal(err)
	}

	queries = repo.New(conn)
}

type SafePlayerRecorder struct {
	mu sync.Mutex
	v  map[string]playermodels.Player
}

func (s *SafePlayerRecorder) Insert(p playermodels.Player) {
	s.mu.Lock()
	if _, ok := s.v[p.PolarisID]; !ok {
		s.v[p.PolarisID] = p
	}
	s.mu.Unlock()
}

func (s *SafePlayerRecorder) Length() int {
	return len(s.v)
}

func (s *SafePlayerRecorder) Debug() {
	for _, p := range s.v {
		fmt.Println("Name: ", p.Name)
		fmt.Println("Rank: ", p.Rank)
	}
}

func main() {
	fmt.Println("Fetching Replays...")
	replays, err := getReplays(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Finished Fetching Replays! Scrape Started")

	scrapeStart := time.Now()

	workerCount := 16
	workerQuota := int(math.Max(math.Floor(float64(len(replays)/workerCount)), 1.0))

	start := 0
	end := workerQuota

	finishedWorker := 0

	safe := SafePlayerRecorder{
		v: make(map[string]playermodels.Player),
	}

	for i := range workerCount {
		if start >= len(replays) {
			finishedWorker++
			continue
		}
		if i == workerCount-1 {
			end = len(replays)
		}
		go func(workerID int, s int, e int) {
			replaySlice := replays[s:e]
			for _, replay := range replaySlice {
				p1 := playermodels.Player{
					PolarisID: replay.P1PolarisID,
					Name:      replay.P1Name,
					Rank:      replay.P1Rank,
					RegionID:  replay.P1RegionID,
					Timestamp: db.Timestamp{
						UpdatedAt: replay.BattleAt,
					},
				}
				p2 := playermodels.Player{
					PolarisID: replay.P2PolarisID,
					Name:      replay.P2Name,
					Rank:      replay.P2Rank,
					RegionID:  replay.P2RegionID,
					Timestamp: db.Timestamp{
						UpdatedAt: replay.BattleAt,
					},
				}

				safe.Insert(p1)
				safe.Insert(p2)
			}

			finishedWorker++
		}(i, start, end)

		start = end
		end += workerQuota
		end = min(end, len(replays))
	}

	for finishedWorker < workerCount {
	}

	fmt.Println("Scraping finished. It took", time.Since(scrapeStart).String())

	// TODO: Should only log a new data if the player has updated their name, rank, and region. and it has to be from a more recent data than the existing one BattleAt > UpdatedAt
	// TODO: Create a player character database to log known character usages

	dbInsertStart := time.Now()
	fmt.Println("Inserting to database...")
	for _, p := range safe.v {
		queries.InsertNewPlayer(context.Background(), repo.InsertNewPlayerParams{
			PolarisID: p.PolarisID,
			// CharaID:   int32(p.CharaID), // We might want to move this to another database
			// Power:     int32(p.Power),    // or this
			Name:      p.Name,            // This should definitely update
			Rank:      int32(p.Rank),     // This should update accordingly
			RegionID:  int32(p.RegionID), // This should as well
			CreatedAt: time.Now().Unix(),
			UpdatedAt: p.UpdatedAt,
		})
	}
	fmt.Println("Inserting finished. It took", time.Since(dbInsertStart).String())
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
