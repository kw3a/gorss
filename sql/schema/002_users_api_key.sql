-- +goose Up
alter table users
add column api_key varchar(64) UNIQUE NOT NULL DEFAULT encode(sha256(random()::text::bytea), 'hex');

-- +goose Down
alter table users
drop column api_key;