-- name: InsertPlayerCharacter :one
insert into player_characters (
	polaris_id, chara_id, created_at, updated_at
) values (
	$1, $2, $3, $4
)
returning *;

-- name: FindPlayerCharacter :one
select * from player_characters where polaris_id = $1 AND chara_id = $2 limit 1;


