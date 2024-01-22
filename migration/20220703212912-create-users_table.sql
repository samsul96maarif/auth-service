/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2022-07-03 21:29:23
 * @modify date 2022-07-03 21:29:23
 * @desc [description]
 */

-- +migrate Up
CREATE TABLE users (
    id serial PRIMARY KEY,
    email varchar(255) UNIQUE NOT NULL,
    name varchar(255),
    password text,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE users;