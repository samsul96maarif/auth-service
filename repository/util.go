/*
 * Author: Samsul Ma'arif <samsulma828@gmail.com>
 * Copyright (c) 2023.
 */

package repository

import (
	"context"

	"github.com/samsul96maarif/auth-service/lib"
)

type RepoUtil interface {
	Transaction(ctx context.Context, fn func(context.Context) error) (err error)
}

type util struct {
	db *lib.Database
}

func NewRepoUtil(db *lib.Database) *util { return &util{db} }

func (repo *util) Transaction(ctx context.Context, fn func(context.Context) error) (err error) {
	trx := repo.db.Begin()
	ctx = context.WithValue(ctx, "Trx", &lib.Database{DB: trx})
	if err = fn(ctx); err != nil {
		trx.Rollback()
	} else {
		err = trx.Commit().Error
	}
	return err
}
