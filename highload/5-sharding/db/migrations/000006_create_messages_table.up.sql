CREATE TABLE messages (
	id        VARCHAR(36) NOT NULL,
    from_user_id VARCHAR(36) NOT NULL,
    to_user_id VARCHAR(36) NOT NULL,
	text     VARCHAR(128) NOT NULL,
    FOREIGN KEY (from_user_id) REFERENCES users (id) ON DELETE CASCADE
    FOREIGN KEY (to_user_id) REFERENCES users (id) ON DELETE CASCADE
	PRIMARY KEY (id)
);