package cheaters

import (
	"context"
	"time"

	"github.com/Lunarisnia/argus-tekken/database/repo"
	"github.com/Lunarisnia/argus-tekken/internal/cheaters/cheaterparams"
)

type CheaterService interface {
	NewCheater(ctx context.Context, newCheater cheaterparams.NewCheater) error
}

type cheaterServiceImpl struct {
	ctx context.Context
	db  *repo.Queries
}

func NewCheaterService(db *repo.Queries) CheaterService {
	return &cheaterServiceImpl{
		db: db,
	}
}

func (ch cheaterServiceImpl) NewCheater(ctx context.Context, newCheater cheaterparams.NewCheater) error {
	_, err := ch.db.InsertNewCheater(ctx, repo.InsertNewCheaterParams{
		PolarisID: newCheater.PolarisID,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	})
	if err != nil {
		return err
	}
	return nil
}
