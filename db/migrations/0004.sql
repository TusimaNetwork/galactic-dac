-- +migrate Down
DROP SCHEMA IF EXISTS near_cache CASCADE;

-- +migrate Up
CREATE SCHEMA near_cache;

CREATE TABLE near_cache.offchain_log
(
    id  SERIAL PRIMARY KEY,
    key VARCHAR UNIQUE
);

CREATE TABLE near_cache.state_log
(
    id  SERIAL PRIMARY KEY,
    log_id INT UNIQUE,
    tx VARCHAR
);