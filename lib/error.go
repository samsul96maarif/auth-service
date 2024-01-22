/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2022-07-03 15:58:51
 * @modify date 2022-07-03 15:58:51
 * @desc [description]
 */
package lib

import "net/http"

type CustomError struct {
	Message  string
	Field    string
	Code     int
	HttpCode int
}

func (err CustomError) Error() string {
	return err.Message
}

func CustomInternalServerError(field, message string, code int) error {
	if code == 0 {
		code = 1011
	}
	return CustomError{
		Message:  message,
		Field:    field,
		Code:     code,
		HttpCode: http.StatusInternalServerError,
	}
}

func InvalidParameterError(field, message string) (err error) {
	err = CustomError{
		Message:  message,
		Field:    field,
		Code:     422,
		HttpCode: http.StatusUnprocessableEntity,
	}
	return err
}

var (
	ErrorUnauthorized = CustomError{
		Message:  "Unauthorized",
		Code:     1000,
		HttpCode: http.StatusUnauthorized,
	}

	ErrorForbidden = CustomError{
		Message:  "Forbidden",
		Code:     1001,
		HttpCode: http.StatusForbidden,
	}

	ErrorInvalidParameter = CustomError{
		Message:  "Invalid Parameter",
		Code:     1002,
		HttpCode: http.StatusUnprocessableEntity,
	}

	ErrorNotFound = CustomError{
		Message:  "Not Found",
		Code:     1003,
		HttpCode: http.StatusNotFound,
	}

	ErrorCannotDeleteDefaultPrice = CustomError{
		Message:  "Cannot Delete Default Price or Capital Price",
		Code:     1004,
		HttpCode: http.StatusUnprocessableEntity,
	}

	ErrorInternalServer = CustomError{
		Message:  "Something went wrong",
		Code:     1011,
		HttpCode: http.StatusInternalServerError,
	}
)
