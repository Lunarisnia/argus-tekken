create table players (
	id serial primary key,
	polaris_id text not null,
	name text not null,
	rank int not null,
	region_id int not null,
	created_at bigint not null,
	updated_at bigint not null
)
