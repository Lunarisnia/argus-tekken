package players

import (
	"context"
	"errors"
	"time"

	"github.com/Lunarisnia/argus-tekken/database/repo"
	"github.com/Lunarisnia/argus-tekken/internal/players/playermodels"
	"github.com/jackc/pgx/v5"
)

type PlayerService interface {
	InsertNewPlayer(ctx context.Context, newPlayer playermodels.Player) error
}

type playerServiceImpl struct {
	db *repo.Queries
}

func NewPlayerService(q *repo.Queries) PlayerService {
	return &playerServiceImpl{
		db: q,
	}
}

func (p playerServiceImpl) InsertNewPlayer(ctx context.Context, newPlayer playermodels.Player) error {
	playerInfo, err := p.db.FindLatestPlayerInfoByPolarisID(ctx, newPlayer.PolarisID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	if ((playerInfo.Name != newPlayer.Name) || (playerInfo.Rank != newPlayer.Rank) || (playerInfo.RegionID != newPlayer.RegionID)) && (newPlayer.UpdatedAt > playerInfo.UpdatedAt) {
		_, err := p.db.InsertNewPlayer(ctx, repo.InsertNewPlayerParams{
			// CharaID:   int32(p.CharaID), // We might want to move this to another database
			// Power:     int32(p.Power),    // or this
			PolarisID: newPlayer.PolarisID,
			Name:      newPlayer.Name,            // This should definitely update
			Rank:      int32(newPlayer.Rank),     // This should update accordingly
			RegionID:  int32(newPlayer.RegionID), // This should as well
			CreatedAt: time.Now().Unix(),
			UpdatedAt: newPlayer.UpdatedAt,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
