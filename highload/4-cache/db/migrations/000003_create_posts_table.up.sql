CREATE TABLE posts(
	id        VARCHAR(36) NOT NULL,
    author_id VARCHAR(36) NOT NULL,
    text VARCHAR(1024) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (author_id) REFERENCES users (id) ON DELETE CASCADE
);