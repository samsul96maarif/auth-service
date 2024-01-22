/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2022-10-21 11:36:28
 * @modify date 2022-10-21 11:36:28
 * @desc [description]
 */

-- +migrate Up
CREATE TABLE roles (
    id serial PRIMARY KEY,
    role varchar(255) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO roles (id, role) VALUES (1, 'SUPER_ADMIN'), (2, 'ADMIN'), (3, 'STAFF');

-- +migrate Down
DROP TABLE roles;