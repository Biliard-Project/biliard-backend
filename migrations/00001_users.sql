-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email TEXT UNIQUE NOT NULL,
  name TEXT NOT NULL,
  password TEXT NOT NULL
);

INSERT INTO users (email, name, password)
VALUES ('test@test.com', 'Admin', '$2a$10$K1UlEymIpjwMVg5B83f.Keattinn.hFdLjJwqt0mJWkQ44Z18IfQW');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
