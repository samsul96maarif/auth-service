/*
 * Author: Samsul Ma'arif<samsulma828@gmail.com>
 * Copyright (c) 2024.
 */

package route

import (
	"net/http"
)

func (route ApiRoute) HelloRoute() {
	route.R.HandleFunc("/hello", route.Handler.AuthMiddleware(http.HandlerFunc(route.Handler.GetHello), "hello").ServeHTTP).Methods(http.MethodGet)
	route.R.HandleFunc("/hello", route.Handler.AuthMiddleware(http.HandlerFunc(route.Handler.AddHello), "add-hello").ServeHTTP).Methods(http.MethodPost)
	route.R.HandleFunc("/Hello/{id}", route.Handler.AuthMiddleware(http.HandlerFunc(route.Handler.EditHello), "edit-hello").ServeHTTP).Methods(http.MethodPut)
	route.R.HandleFunc("/Hello/{id}", route.Handler.AuthMiddleware(http.HandlerFunc(route.Handler.DeleteHello), "delete-hello").ServeHTTP).Methods(http.MethodDelete)
}
