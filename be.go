/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2023-09-09 18:26:15
 * @modify date 2023-09-09 18:26:15
 * @desc [description]
 */
package goapiapp

import (
	"github.com/samsul96maarif/auth-service/lib"
	"github.com/samsul96maarif/auth-service/usecase"
)

type BE struct {
	Auth    usecase.Auth
	Usecase *usecase.Usecase
}

func NewBe(db *lib.Database) (api BE) {
	u := usecase.NewUsecase(db)
	api.Usecase = &u
	return
}

func (app BE) Init() {

}
