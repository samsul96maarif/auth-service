/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2022-07-03 20:58:28
 * @modify date 2022-07-03 20:58:28
 * @desc [description]
 */
package route

import (
	"github.com/samsul96maarif/auth-service/handler"

	"github.com/gorilla/mux"
)

type ApiRoute struct {
	R       *mux.Router
	Handler *handler.Handler
}
