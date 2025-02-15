package charamodels

import "github.com/Lunarisnia/argus-tekken/internal/db"

type Character struct {
	ID          uint   `json:"id"`
	PolarisID   string `json:"polaris_id"`
	CharacterID int    `json:"character_id"`
	db.Timestamp
}
