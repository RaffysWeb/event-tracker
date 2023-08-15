CREATE TABLE IF NOT EXISTS mykeyspace.events_by_trackable (
	id	        				UUID,
	actor_id 						TEXT,
	trackable_owner_id  TEXT,
	event_type					TEXT,
	trackable_type 			TEXT,
	trackable_id 				TEXT,
	created_at       		TIMESTAMP,
	metadata  					MAP<TEXT, TEXT>,

	PRIMARY KEY ((trackable_type, trackable_id), created_at)
);