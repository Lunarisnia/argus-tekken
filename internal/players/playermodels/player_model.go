package playermodels

import "github.com/Lunarisnia/argus-tekken/internal/db"

type Player struct {
	ID        uint   `json:"id"`
	PolarisID string `json:"polaris_id"`
	Name      string `json:"name"`
	Rank      int32  `json:"rank"`
	RegionID  int32  `json:"region_id"`
	db.Timestamp
}
