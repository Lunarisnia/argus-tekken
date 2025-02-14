package playermodels

import "github.com/Lunarisnia/argus-tekken/internal/db"

type Player struct {
	ID        uint   `json:"id"`
	PolarisID string `json:"polaris_id"`
	Name      string `json:"name"`
	Power     int    `json:"power"`
	Rank      int    `json:"rank"`
	RegionID  int    `json:"region_id"`
	db.Timestamp
}
