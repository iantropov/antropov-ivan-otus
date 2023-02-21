CREATE TABLE users_dialogs(
    user_id_1 VARCHAR(36) NOT NULL,
    user_id_2 VARCHAR(36) NOT NULL,
    dialog_id VARCHAR(36) NOT NULL,
    PRIMARY KEY (user_id_1, user_id_2),
    FOREIGN KEY (user_id_1)  REFERENCES users  (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id_2)  REFERENCES users  (id) ON DELETE CASCADE,
    FOREIGN KEY (dialog_id) REFERENCES dialogs (id) ON DELETE CASCADE
);