-- name: InsertNewPlayer :one
insert into players (
	id, polaris_id, name, power, rank, region_id, created_at, updated_at
) values (
	$1, $2,	$3, $4,	$5, $6, $7, $8
)
returning *;

