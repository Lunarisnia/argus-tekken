package wankmodels

// {
// 	"battle_at": 1739429575,

// 	"p1_chara_id": 9, // Maybe
// 	"p1_name": ".......",
// 	"p1_polaris_id": "4qjMy5mTjfEr",
// 	"p1_power": 327334,
// 	"p1_rank": 27,
// 	"p1_region_id": 0,

// 	"p1_user_id": 260015190910102553,
// 	"p2_chara_id": 4, // Maybe
// 	"p2_name": "SSJ_H",
// 	"p2_polaris_id": "35THMbb4M67B",
// 	"p2_power": 249517,
// 	"p2_rank": 26,
// 	"p2_region_id": 0,
// },

type Replay struct {
	// Metadata
	BattleAt int64 `json:"battle_at"`

	// Player 1
	P1CharaID   int    `json:"p1_chara_id"`
	P1Name      string `json:"p1_name"`
	P1PolarisID string `json:"p1_polaris_id"`
	P1Power     int    `json:"p1_power"`
	P1Rank      int    `json:"p1_rank"`
	P1RegionID  int    `json:"p1_region_id"`

	// Player 2
	P2CharaID   int    `json:"p2_chara_id"`
	P2Name      string `json:"p2_name"`
	P2PolarisID string `json:"p2_polaris_id"`
	P2Power     int    `json:"p2_power"`
	P2Rank      int    `json:"p2_rank"`
	P2RegionID  int    `json:"p2_region_id"`
}
