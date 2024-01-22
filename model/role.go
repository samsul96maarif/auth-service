/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2022-10-22 08:07:14
 * @modify date 2022-10-22 08:07:14
 * @desc [description]
 */

package model

type Role struct {
	BaseModel
	Role string `json:"role"`
}

type UserRole struct {
	RoleId uint `json:"role_id" gorm:"primaryKey"`
	UserId uint `json:"user_id" gorm:"primaryKey"`
}

type RolePermission struct {
	PermissionSlug string `json:"permission_role" gorm:"primaryKey"`
	RoleId         uint   `json:"role_id" gorm:"primaryKey"`
}
