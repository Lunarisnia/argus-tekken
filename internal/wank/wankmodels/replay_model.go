package wankmodels

type Replay struct {
	// Metadata
	BattleAt int64 `json:"battle_at"`

	// Player 1
	P1CharaID   int32  `json:"p1_chara_id"`
	P1Name      string `json:"p1_name"`
	P1PolarisID string `json:"p1_polaris_id"`
	P1Power     int    `json:"p1_power"`
	P1Rank      int32  `json:"p1_rank"`
	P1RegionID  int32  `json:"p1_region_id"`

	// Player 2
	P2CharaID   int32  `json:"p2_chara_id"`
	P2Name      string `json:"p2_name"`
	P2PolarisID string `json:"p2_polaris_id"`
	P2Power     int    `json:"p2_power"`
	P2Rank      int32  `json:"p2_rank"`
	P2RegionID  int32  `json:"p2_region_id"`
}
