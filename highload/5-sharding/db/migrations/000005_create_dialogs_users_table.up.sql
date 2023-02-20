CREATE TABLE dialogs_users(
    dialog_id VARCHAR(36) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    PRIMARY KEY (dialog_id, user_id),
    FOREIGN KEY (dialog_id) REFERENCES dialogs (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id)  REFERENCES users  (id) ON DELETE CASCADE
);