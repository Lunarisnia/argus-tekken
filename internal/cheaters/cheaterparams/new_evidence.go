package cheaterparams

type NewEvidence struct {
	PolarisID       string `json:"polaris_id"`
	OffenceCategory string `json:"offence_category"`
	EvidenceURL     string `json:"evidence_url"`
}
