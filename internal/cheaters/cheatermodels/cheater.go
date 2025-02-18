package cheatermodels

import "github.com/Lunarisnia/argus-tekken/internal/db"

type Cheater struct {
	ID        uint   `json:"id"`
	PolarisID string `json:"polaris_id"`
	db.Timestamp
}
