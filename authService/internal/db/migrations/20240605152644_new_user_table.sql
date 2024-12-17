-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL primary key ,
                                     user_id varchar(200) unique,
    username varchar(255) NOT NULL,
    password varchar(255) NOT NULL
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
