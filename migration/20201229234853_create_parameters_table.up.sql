CREATE TABLE parameters (
    user_id INTEGER NOT NULL,
    secret TEXT NOT NULL,
    issuer TEXT NOT NULL,
    account TEXT NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id)
		ON DELETE CASCADE ON UPDATE CASCADE,
	PRIMARY KEY (user_id, secret)
)
