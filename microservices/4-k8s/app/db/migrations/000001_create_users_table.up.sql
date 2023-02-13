CREATE TABLE users (
    id  BIGINT NOT NULL AUTO_INCREMENT,
	username      VARCHAR(128) NOT NULL,
	first_name      VARCHAR(128) NOT NULL,
	second_name     VARCHAR(128) NOT NULL,
	email     VARCHAR(128) NOT NULL,
	phone     VARCHAR(128) NOT NULL,
	PRIMARY KEY (`id`)
);