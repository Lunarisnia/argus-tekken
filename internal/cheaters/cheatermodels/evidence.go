package cheatermodels

import "github.com/Lunarisnia/argus-tekken/internal/db"

type Evidence struct {
	ID              uint   `json:"id"`
	PolarisID       string `json:"polaris_id"`
	EvidenceURL     string `json:"evidence_url"`
	OffenceCategory string `json:"offence_category"`
	db.Timestamp
}
