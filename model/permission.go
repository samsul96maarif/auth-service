/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2024-01-20 10:23:51
 * @modify date 2024-01-20 10:23:51
 * @desc [description]
 */
package model

import "time"

type Permission struct {
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	HttpMethod string    `json:"http_method"`
	Slug       string    `json:"slug" gorm:"primaryKey"`
	Url        string    `json:"url"`
}
