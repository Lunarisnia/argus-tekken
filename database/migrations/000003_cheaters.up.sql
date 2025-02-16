create table cheaters (
	id serial primary key,
	polaris_id text unique not null,
	created_at bigint not null,
	updated_at bigint not null
)
