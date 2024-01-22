/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2022-07-03 20:52:18
 * @modify date 2022-07-03 20:52:18
 * @desc [description]
 */
package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	goapiapp "github.com/samsul96maarif/auth-service"
	"github.com/samsul96maarif/auth-service/lib"
	"github.com/samsul96maarif/auth-service/request"
)

type RequestValidator interface {
	Ok() error
}

func Decode(r *http.Request, v RequestValidator) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return lib.InvalidParameterError("", err.Error())
	}
	return v.Ok()
}

type Handler struct{ BE *goapiapp.BE }

type ResponseMeta struct {
	Page       *uint `json:"page,omitempty"`
	PerPage    *uint `json:"per_page,omitempty"`
	Total      *uint `json:"total,omitempty"`
	HttpStatus int   `json:"http_status"`
}

type ResponseBody struct {
	Meta    ResponseMeta `json:"meta"`
	Data    interface{}  `json:"data,omitempty"`
	Message string       `json:"message,omitempty"`
}

type ErrorInfo struct {
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
	Code    int    `json:"code"`
}

type ErrorBody struct {
	Errors []ErrorInfo `json:"errors"`
	Meta   interface{} `json:"meta"`
}

func GenerateListRequest(params url.Values) request.ListRequest {
	return request.ListRequest{strings.Split(params.Get("sort"), ","), params.Get("search")}
}

func GenerateBaseRequest(r *http.Request) request.BaseRequest {
	perPageInt, _ := strconv.ParseUint(r.URL.Query().Get("per_page"), 10, 64)
	pageInt, _ := strconv.ParseUint(r.URL.Query().Get("page"), 10, 64)
	return request.BaseRequest{
		Page:  uint(pageInt),
		Limit: uint(perPageInt),
	}
}

func NewHandler(be *goapiapp.BE) Handler {
	return Handler{BE: be}
}

func writeError(w http.ResponseWriter, err error) {
	var resp interface{}
	code := http.StatusInternalServerError

	switch errOrig := err.(type) {
	case lib.CustomError:
		resp = ErrorBody{
			Errors: []ErrorInfo{
				{
					Message: errOrig.Message,
					Code:    errOrig.Code,
					Field:   errOrig.Field,
				},
			},
			Meta: ResponseMeta{
				HttpStatus: errOrig.HttpCode,
			},
		}

		code = int(errOrig.HttpCode)
	default:
		resp = ResponseBody{
			Message: "Internal Server Error",
			Meta: ResponseMeta{
				HttpStatus: code,
			},
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

func writeSuccess(w http.ResponseWriter, data interface{}, message string, meta ResponseMeta) {
	resp := ResponseBody{
		Message: message,
		Data:    data,
		Meta:    meta,
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(meta.HttpStatus)
	json.NewEncoder(w).Encode(resp)
}

func writeResponse(w http.ResponseWriter, resp interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}
