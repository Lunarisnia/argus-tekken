package cheaters

import (
	"context"

	"github.com/Lunarisnia/argus-tekken/database/repo"
	"github.com/Lunarisnia/argus-tekken/internal/cheaters/cheatermodels"
)

type CheaterService interface {
	NewCheater(ctx context.Context, newCheater cheatermodels.Cheater) error
}

type cheaterServiceImpl struct {
	db *repo.Queries
}

func NewCheaterService(db *repo.Queries) CheaterService {
	return &cheaterServiceImpl{
		db: db,
	}
}

func (ch cheaterServiceImpl) NewCheater(ctx context.Context, newCheater cheatermodels.Cheater) error {
	// TODO: Create db query to insert new cheater
	return nil
}
