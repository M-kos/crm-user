-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  uuid VARCHAR(255) NOT NULL UNIQUE,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  user_type VARCHAR(255) NOT NULL,
  phone VARCHAR(255),
  tg VARCHAR(255),
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_uuid ON users (uuid);
CREATE INDEX idx_users_email ON users (email);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
