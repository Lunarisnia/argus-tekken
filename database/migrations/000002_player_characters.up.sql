create table player_characters (
	id serial primary key,
	polaris_id text not null,
	chara_id int not null,
	created_at bigint not null,
	updated_at bigint not null
)
