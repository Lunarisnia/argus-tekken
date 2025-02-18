-- name: InsertNewEvidence :one
insert into evidences (
	polaris_id, evidence_url, offence_category, created_at, updated_at
) values (
	$1, $2, $3, $4, $5
) returning *;
