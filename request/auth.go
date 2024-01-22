/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2024-01-20 10:13:17
 * @modify date 2024-01-20 10:13:17
 * @desc [description]
 */
package request

import (
	"net/http"

	"github.com/samsul96maarif/auth-service/lib"
)

type ForgotPassword struct {
	Email string `json:"email" validate:"required"`
}

type Login struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Register struct {
	Name                 string `json:"name" validate:"required"`
	Email                string `json:"email" validate:"required"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required"`
}

func (r *Register) Ok() error {
	if r.Password != r.PasswordConfirmation {
		return lib.CustomError{
			Message:  "password confirmation doesnt match",
			Field:    "password_confirmation",
			HttpCode: http.StatusBadRequest,
		}
	}
	return nil
}
