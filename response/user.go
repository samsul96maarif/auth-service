/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2024-01-20 10:28:24
 * @modify date 2024-01-20 10:28:24
 * @desc [description]
 */
package response

import (
	"time"

	"github.com/samsul96maarif/auth-service/model"
)

type User struct {
	Roles     []Role    `json:"roles,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Id        uint      `json:"id,omitempty"`
}

func NewUser(entity model.User, roles []model.Role) (res User) {
	var rolesRes []Role
	for _, role := range roles {
		rolesRes = append(rolesRes, NewRole(role))
	}
	res = User{
		Roles:     rolesRes,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		Name:      entity.Name,
		Email:     entity.Email,
		Id:        entity.Id,
	}
	return
}
