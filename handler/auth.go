/*
 * Author: Samsul Ma'arif<samsulma828@gmail.com>
 * Copyright (c) 2024.
 */

package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/samsul96maarif/auth-service/lib"
	"github.com/samsul96maarif/auth-service/request"
)

func (handler *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var payload request.Login
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		writeError(w, lib.InvalidParameterError("", err.Error()))
		return
	}
	err = validator.New().Struct(payload)
	if err != nil {
		writeError(w, lib.InvalidParameterError("", err.Error()))
		return
	}
	res, err := handler.BE.Auth.Login(r.Context(), payload)
	if err != nil {
		writeError(w, err)
		return
	}
	writeSuccess(w, res, "Success", ResponseMeta{HttpStatus: http.StatusOK})
	return
}
