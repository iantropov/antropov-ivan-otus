CREATE TABLE friends(
    user_id INT NOT NULL,
    friend_id INT NOT NULL,
    PRIMARY KEY (user_id, friend_id),
    FOREIGN KEY user_id references users (id) on delete cascade,
    FOREIGN KEY friend_id  references users  (id) on delete cascade
}