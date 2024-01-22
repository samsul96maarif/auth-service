/*
 * Author: Samsul Ma'arfi <samsulma828@gmail.com>
 * Copyright (c) 2022.
 */

package request

import (
	"fmt"
	"mime/multipart"
	"strings"
)

type ListReq interface {
	GetLikeColumns() []string
	GetKeyword() string
	GetOrderQuery() string
}

type PaginateReq interface {
	ListReq
	GetPage() uint
	GetLimit() uint
	GetOffset() int
}

type ListRequest struct {
	Sorts   []string `json:"sorts,omitempty;query:sorts"`
	Keyword string   `json:"keyword,omitempty;query:keyword"`
}

func (req *ListRequest) GetKeyword() string {
	return req.Keyword
}

type BaseRequest struct {
	Page  uint `json:"page,omitempty;query:page"`
	Limit uint `json:"limit,omitempty;query:limit"`
}

func (req *BaseRequest) GetOffset() int {
	page := req.GetPage()
	limit := req.GetLimit()
	offset := int((page - 1) * limit)
	return offset
}

func (req *BaseRequest) GetPage() uint {
	if req.Page > 0 {
		return req.Page
	}
	return 1
}

func (req *BaseRequest) GetLimit() uint {
	if req.Limit > 0 {
		return req.Limit
	}
	return 100
}

type BaseUploadFileRequest struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
}

type FindRequest interface {
	GetIdentifier() string
	GetOrderQuery() string
}

type findRequest struct {
	Sorts      []string
	Identifier string
}

func NewFindRequest(sorts []string, identifier string) *findRequest {
	return &findRequest{sorts, identifier}
}

func (r *findRequest) GetIdentifier() string {
	return r.Identifier
}

func (r *findRequest) GetOrderQuery() string {
	var result []string
	fieldMap := map[string]string{
		"created_at": "created_at",
		"updated_at": "updated_at",
	}

	for _, s := range r.Sorts {
		if len(s) == 0 {
			continue
		}

		order, key := "ASC", s
		if s[:1] == "-" {
			order, key = "DESC", s[1:]
		}

		fieldName, ok := fieldMap[key]
		if !ok {
			continue
		}

		result = append(result, fmt.Sprintf("%s %s", fieldName, order))
	}

	return strings.Join(result, ",")
}

type StoreRequest interface {
	GetData() interface{}
}

type storeRequest struct {
	data interface{}
}

func NewStoreRequest(data interface{}) *storeRequest { return &storeRequest{data} }

func (r *storeRequest) GetData() interface{} { return r.data }

type UpdateRequest interface {
	GetData() interface{}
	GetWhere() map[string]interface{}
}

type updateRequest struct {
	where map[string]interface{}
	data  interface{}
}

func NewUpdateRequest(where map[string]interface{}, data interface{}) *updateRequest {
	return &updateRequest{where, data}
}

func (r *updateRequest) GetData() interface{} { return r.data }

func (r *updateRequest) GetWhere() map[string]interface{} { return r.where }
