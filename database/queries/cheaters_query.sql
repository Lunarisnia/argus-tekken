-- name: InsertNewCheater :one
insert into cheaters (
	polaris_id, created_at, updated_at
) values ($1, $2, $3)
returning *;
