CREATE TABLE users (
	id        VARCHAR(36) NOT NULL,
	first_name      VARCHAR(128) NOT NULL,
	second_name     VARCHAR(128) NOT NULL,
	age     INT NOT NULL,
	password     VARCHAR(128),
	biography     VARCHAR(255),
	city     VARCHAR(64),
	PRIMARY KEY (`id`)
);