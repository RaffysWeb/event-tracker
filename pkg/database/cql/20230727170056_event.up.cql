CREATE TABLE IF NOT EXISTS mykeyspace.events (
	id        TEXT PRIMARY KEY,
	eventType TEXT,
	userId    TEXT,
	metadata  MAP<TEXT, TEXT>
);