/*
 * Author: Samsul Ma'arif <samsulma828@gmail.com>
 * Copyright (c) 2023.
 */

package model

import "time"

type BaseModel struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WebsocketMessage struct {
	Message     interface{} `json:"message"`
	CommandType string      `json:"command_type"    validate:"required"`
	Language    string      `json:"language"        options:"id,en"`
	Token       string      `json:"token"`
	From        string      `json:"from"`
	To          string      `json:"to"`
}
