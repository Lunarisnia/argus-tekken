// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package repo

type Cheater struct {
	ID        int32
	PolarisID string
	CreatedAt int64
	UpdatedAt int64
}

type Evidence struct {
	ID              int32
	PolarisID       string
	EvidenceUrl     string
	OffenceCategory string
	CreatedAt       int64
	UpdatedAt       int64
}

type Player struct {
	ID        int32
	PolarisID string
	Name      string
	Rank      int32
	RegionID  int32
	CreatedAt int64
	UpdatedAt int64
}

type PlayerCharacter struct {
	ID        int32
	PolarisID string
	CharaID   int32
	CreatedAt int64
	UpdatedAt int64
}
