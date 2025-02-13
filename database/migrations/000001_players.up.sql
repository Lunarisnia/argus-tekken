-- TODO: Verify this is working by setting up docker and database
create table players (
	id serial primary key,
	polaris_id string,
	name text,
	power int,
	rank int,
	region_id int,
	created_at timestampz,
	updated_at timestampz
)
