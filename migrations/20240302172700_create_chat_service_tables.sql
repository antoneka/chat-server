-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id INT UNIQUE NOT NULL CHECK (id > 0),
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS chats (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    creator_id INT NOT NULL REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    message TEXT NOT NULL,
    chat_id INT NOT NULL REFERENCES chats(id),
    sender_id INT NOT NULL REFERENCES users(id),
    timestamp TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS chat_members (
    chat_id INT NOT NULL REFERENCES chats(id),
    user_id INT NOT NULL REFERENCES users(id),
    PRIMARY KEY (chat_id, user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS chats;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS chat_users;
-- +goose StatementEnd