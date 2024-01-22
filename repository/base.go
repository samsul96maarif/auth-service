/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2022-07-03 13:56:24
 * @modify date 2022-07-03 13:56:24
 * @desc [description]
 */
/*
 * Author: Samsul Ma'arif <samsulma828@gmail.com>
 * Copyright (c) 2022.
 */

package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/samsul96maarif/auth-service/lib"
	"github.com/samsul96maarif/auth-service/lib/logger"
	"github.com/samsul96maarif/auth-service/model"

	"gorm.io/gorm"
)

type BaseRepo interface {
	Delete(ctx context.Context, where map[string]interface{}) error
	GetTableName() string
}

type CountRepo interface {
	Count(ctx context.Context, opts ...OptRepo) (total int64, err error)
}

type OptRepo func(*gorm.DB)

func OptWithLimit(data int) OptRepo {
	return func(stmt *gorm.DB) {
		stmt = stmt.Limit(data)
	}
}

func OptWithKeyword(columns []string, q string) OptRepo {
	return func(stmt *gorm.DB) {
		if q == "" {
			return
		}
		var (
			args []interface{}
			cols []string
		)
		for _, col := range columns {
			cols = append(cols, fmt.Sprintf("%s LIKE ?", col))
			args = append(args, fmt.Sprintf("%%%s%%", q))
		}
		query := strings.Join(cols, " OR ")
		stmt = stmt.Where(query, args...)
	}
}

func OptWithOffset(data int) OptRepo {
	return func(stmt *gorm.DB) {
		stmt = stmt.Offset(data)
	}
}

func OptWithOrder(data string) OptRepo {
	return func(stmt *gorm.DB) {
		stmt = stmt.Order(data)
	}
}

func OptWithSearch(columns []string, search string) OptRepo {
	return func(stmt *gorm.DB) {
		if search == "" {
			return
		}
		var (
			args       []interface{}
			conditions []string
		)
		for _, col := range columns {
			args = append(args, fmt.Sprintf("%%%s%%", search))
			conditions = append(conditions, fmt.Sprintf("%s LIKE ?", col))
		}
		query := strings.Join(conditions, " OR ")
		stmt = stmt.Where(query, args...)
		return
	}
}

func OptWithSelectColumn(data []string) OptRepo {
	return func(stmt *gorm.DB) {
		stmt = stmt.Select(data)
	}
}

func OptWithWhere(data interface{}) OptRepo {
	return func(stmt *gorm.DB) {
		stmt = stmt.Where(data)
	}
}

type Repository struct {
	db *lib.Database
}

func NewRepository(db *lib.Database) Repository { return Repository{db: db} }

func (repo *Repository) Transaction(ctx context.Context, fn func(context.Context) error) (err error) {
	trx := repo.db.Begin()
	ctx = context.WithValue(ctx, "Trx", &lib.Database{DB: trx})
	if err = fn(ctx); err != nil {
		trx.Rollback()
	} else {
		err = trx.Commit().Error
	}
	return err
}

func (repo *Repository) GetUser(ctx context.Context, where map[string]interface{}, order_by string) (entitis []model.User, err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}
	if order_by == "" {
		order_by = "created_at asc"
	}
	err = tx.Where(where).Order(order_by).Find(&entitis).Error
	if err != nil {
		logger.Error(ctx, "Error GetUser ", map[string]interface{}{
			"error": err,
			"tags":  []string{"repo", "users", "get"},
		})
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return entitis, err
		}
	}
	return entitis, nil
}

func (repo *Repository) CreatePermission(ctx context.Context, entity *model.Permission) error {
	tx := getDBConnection(ctx, repo.db)
	if err := tx.Create(entity).Error; err != nil {
		logger.Error(ctx, "Error Create Permission", map[string]interface{}{
			"error": err,
			"tags":  []string{"repo", "create", "permission"},
		})
		return err
	}
	return nil
}

func (repo *Repository) CreateRolePermission(ctx context.Context, entity *model.RolePermission) error {
	tx := getDBConnection(ctx, repo.db)
	if err := tx.Create(entity).Error; err != nil {
		logger.Error(ctx, "Error Create Role Permission", map[string]interface{}{
			"error": err,
			"tags":  []string{"repo", "create", "role_permission"},
		})
		return err
	}
	return nil
}

func (repo *Repository) CreateUser(ctx context.Context, entity *model.User) (err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}
	err = tx.Create(entity).Error
	if err != nil {
		logger.Error(ctx, "Error create user", map[string]interface{}{
			"error": err,
			"tags":  []string{"repo", "create", "user"},
		})
	}
	return err
}

func (repo *Repository) CreateUserRole(ctx context.Context, entity *model.UserRole) error {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}
	err := tx.Create(entity).Error
	if err != nil {
		logger.Error(ctx, "Error CreateUserRole", map[string]interface{}{
			"error": err,
			"tags":  []string{"repo", "user_roles", "create"},
		})
	}
	return err
}

func (repo *Repository) FindUser(ctx context.Context, where map[string]interface{}, order_by string) (entity model.User, err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err = tx.Where(where).Limit(1).Find(&entity).Order(order_by).Error
	if err != nil {
		logger.Error(ctx, "Error find user", map[string]interface{}{
			"error": err,
			"tags":  []string{"repo", "find", "user"},
		})
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return entity, err
		}
	}
	return entity, nil
}

func (repo *Repository) FindUserRole(ctx context.Context, where map[string]interface{}, order_by string) (entity model.UserRole, err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}
	err = tx.Where(where).Limit(1).Find(&entity).Order(order_by).Error
	if err != nil {
		logger.Error(ctx, "Error get userRole", map[string]interface{}{
			"error": err,
			"tags":  []string{"repo", "find", "userRole"},
		})
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return entity, err
		}
	}
	return entity, nil
}

func (repo *Repository) GetRole(ctx context.Context, where map[string]interface{}, order_by string) (entities []model.Role, err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err = tx.Where(where).Find(&entities).Order(order_by).Error
	if err != nil {
		logger.Error(ctx, "Error get roles", map[string]interface{}{
			"error": err,
			"tags":  []string{"repo", "get", "role"},
		})
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return entities, err
		}
	}
	return entities, nil
}

func (repo *Repository) GetUserRole(ctx context.Context, where map[string]interface{}, order_by string) (entities []model.UserRole, err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err = tx.Where(where).Find(&entities).Order(order_by).Error
	if err != nil {
		logger.Error(ctx, "Error get user_roles", map[string]interface{}{
			"error": err,
			"tags":  []string{"repo", "get", "role"},
		})
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return entities, err
		}
	}
	return entities, nil
}

func getDBConnection(ctx context.Context, repo_db *lib.Database) *lib.Database {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo_db
	}
	return tx
}
