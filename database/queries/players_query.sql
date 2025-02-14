-- name: InsertNewPlayer :one
insert into players (
	polaris_id, name, rank, region_id, created_at, updated_at
) values (
	$1,	$2, $3,	$4, $5, $6, $7, $8
)
returning *;

