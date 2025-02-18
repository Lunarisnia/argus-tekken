create table evidences (
	id serial primary key,
	polaris_id text references cheaters (polaris_id) not null,
	evidence_url text not null default '',
	offence_category text not null,
	created_at bigint not null,
	updated_at bigint not null
)
