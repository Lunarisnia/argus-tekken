package characters

import (
	"context"
	"errors"
	"time"

	"github.com/Lunarisnia/argus-tekken/database/repo"
	"github.com/jackc/pgx/v5"
)

type CharacterService interface {
	RegisterNewPlayerCharacter(ctx context.Context, polarisId string, charaId int32) error
}

type characterServiceImpl struct {
	db *repo.Queries
}

func NewCharacterService(db *repo.Queries) CharacterService {
	return &characterServiceImpl{
		db: db,
	}
}

func (c characterServiceImpl) RegisterNewPlayerCharacter(ctx context.Context, polarisId string, charaId int32) error {
	playerChara, err := c.db.FindPlayerCharacter(ctx, repo.FindPlayerCharacterParams{
		PolarisID: polarisId,
		CharaID:   charaId,
	})
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	if playerChara.PolarisID == polarisId && playerChara.CharaID == charaId {
		// Character already on the list
		return nil
	}

	_, err = c.db.InsertPlayerCharacter(ctx, repo.InsertPlayerCharacterParams{
		PolarisID: polarisId,
		CharaID:   charaId,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	})
	if err != nil {
		return err
	}

	return nil
}
