/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2024-01-18 16:53:18
 * @modify date 2024-01-18 16:53:18
 * @desc [description]
 */

-- +migrate Up
CREATE TABLE IF NOT EXISTS role_permissions (
    role_id INT NOT NULL,
    permission_slug VARCHAR NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (role_id, permission_slug)
);

INSERT INTO role_permissions (role_id, permission_slug) VALUES (1, '*');

-- +migrate Down
DROP TABLE role_permissions;