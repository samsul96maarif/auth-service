/*
 * Author: Samsul Ma'arif <samsulma828@gmail.com>
 * Copyright (c) 2023.
 */

package response

import (
	"time"

	"github.com/samsul96maarif/auth-service/model"
)

type Role struct {
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Role      string    `json:"role"`
	Id        uint      `json:"id,omitempty"`
}

func NewRole(entity model.Role) Role {
	return Role{
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		Role:      entity.Role,
		Id:        entity.Id,
	}
}
