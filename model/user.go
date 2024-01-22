/*
 * Author: Samsul Ma'arif <samsulma828@gmail.com>
 * Copyright (c) 2023.
 */

package model

type User struct {
	BaseModel
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
