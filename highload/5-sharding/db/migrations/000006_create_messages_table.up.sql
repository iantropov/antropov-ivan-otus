CREATE TABLE messages (
    id INT AUTO_INCREMENT PRIMARY KEY,
    from_user_id VARCHAR(36) NOT NULL,
    to_user_id VARCHAR(36) NOT NULL,
    dialog_id  VARCHAR(36) NOT NULL,
	text     VARCHAR(128) NOT NULL
);
CREATE INDEX messages_dialog_id_idx on messages(dialog_id) using BTREE;
