-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_information (
    id SERIAL primary key,
    user_id varchar(200) unique,
    surname varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    middle_name varchar(255) NOT NULL,
    age int NOT NULL,
    phone_number varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    phone_confirmation boolean default false,
    email_confirmation boolean default false,
    foreign key (user_id) REFERENCES users(user_id) ON DELETE CASCADE
    );
CREATE UNIQUE INDEX  IF NOT EXISTS phone_number ON user_information(phone_number);
CREATE UNIQUE INDEX IF NOT EXISTS email ON user_information(email);
CREATE INDEX IF NOT EXISTS surname ON user_information(surname);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_information;
-- +goose StatementEnd
