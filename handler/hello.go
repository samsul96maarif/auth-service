/*
 * Author: Samsul Ma'arif<samsulma828@gmail.com>
 * Copyright (c) 2024.
 */

package handler

import (
	"fmt"
	"net/http"
)

func (handler *Handler) GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello home")
	return
}

func (handler *Handler) AddHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Suceed Add Hello")
	return
}

func (handler *Handler) EditHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Suceed Edit Hello")
	return
}

func (handler *Handler) DeleteHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Suceed Delete Hello")
	return
}
