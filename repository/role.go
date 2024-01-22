package repository

import (
	"context"
	"errors"

	"github.com/samsul96maarif/auth-service/lib"
	"github.com/samsul96maarif/auth-service/lib/logger"
	"github.com/samsul96maarif/auth-service/model"
	"gorm.io/gorm"
)

type Role interface {
	BaseRepo
	CountRepo
	Find(ctx context.Context, order string, opts ...OptRepo) (*model.Role, error)
	Get(ctx context.Context, opts ...OptRepo) ([]model.Role, error)
}

type role struct {
	*util
	db *lib.Database
}

func NewRoleRepo(db *lib.Database) Role { return &role{NewRepoUtil(db), db} }

func (repo *role) Count(ctx context.Context, opts ...OptRepo) (total int64, err error) {
	tx := getDBConnection(ctx, repo.db)
	stmt := tx.Model(&model.Role{})
	for _, opt := range opts {
		opt(stmt)
	}
	err = stmt.Count(&total).Error
	if err != nil {
		logger.Error(ctx, "error count on"+repo.GetTableName(), map[string]interface{}{
			"error": err.Error(),
			"tags":  []string{"count", repo.GetTableName()},
		})
	}
	return
}

func (repo *role) Delete(ctx context.Context, where map[string]interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (repo *role) Find(ctx context.Context, order string, opts ...OptRepo) (entity *model.Role, err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}
	stmt := tx.Model(&model.Role{})
	for _, opt := range opts {
		opt(stmt)
	}
	err = stmt.Order(order).Find(&entity).Error
	if err != nil {
		logger.Error(ctx, "error find "+repo.GetTableName(), map[string]interface{}{
			"error": err.Error(),
			"tags":  []string{"repo", repo.GetTableName()},
		})
		if !errors.Is(gorm.ErrRecordNotFound, err) {
			return
		}
	}
	return entity, nil
}

func (repo *role) GetTableName() string { return "roles" }

func (repo *role) Get(ctx context.Context, opts ...OptRepo) (entities []model.Role, err error) {
	tx := getDBConnection(ctx, repo.db)
	stmt := tx.Model(&model.Role{})
	for _, opt := range opts {
		opt(stmt)
	}
	err = stmt.Find(&entities).Error
	if err != nil {
		logger.Error(ctx, "error list "+repo.GetTableName(), map[string]interface{}{
			"error": err.Error(),
			"tags":  []string{"repo", repo.GetTableName()},
		})
		if !errors.Is(gorm.ErrRecordNotFound, err) {
			return
		}
	}
	return entities, nil
}

func (repo *Repository) FindRolePermission(ctx context.Context, opts ...OptRepo) (entity *model.RolePermission, err error) {
	tx := getDBConnection(ctx, repo.db)
	stmt := tx.Model(&model.RolePermission{})
	for _, opt := range opts {
		opt(stmt)
	}
	err = stmt.Limit(1).Find(&entity).Error
	return
}

func (repo *Repository) DeleteRolePermission(ctx context.Context, role_id uint, permission_slug string) error {
	tx := getDBConnection(ctx, repo.db)
	err := tx.Model(&model.RolePermission{}).Delete(&model.RolePermission{}, map[string]interface{}{
		"role_id":         role_id,
		"permission_slug": permission_slug,
	}).Error
	return err
}
