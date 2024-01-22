/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2024-01-18 16:53:18
 * @modify date 2024-01-18 16:53:18
 * @desc [description]
 */

-- +migrate Up
CREATE TABLE IF NOT EXISTS permissions (
    slug VARCHAR PRIMARY KEY,
    url VARCHAR,
    http_method VARCHAR,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO permissions (slug, url, http_method) VALUES ('*', '*', '*');

-- +migrate Down
DROP TABLE permissions;