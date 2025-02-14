-- name: InsertNewPlayer :one
insert into players (
	polaris_id, name, rank, region_id, created_at, updated_at
) values (
	$1,	$2, $3,	$4, $5, $6
)
returning *;


-- name: FindLatestPlayerInfoByPolarisID :one
select * from players where polaris_id = $1 order by updated_at desc limit 1;
