/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2022-07-03 15:33:30
 * @modify date 2022-07-03 15:33:30
 * @desc [description]
 */
package repository

import (
	"context"
	"errors"

	"github.com/samsul96maarif/auth-service/lib"
	"github.com/samsul96maarif/auth-service/lib/logger"
	"github.com/samsul96maarif/auth-service/model"

	"gorm.io/gorm"
)

type User interface {
	BaseRepo
	CountRepo
	RepoUtil
	Find(ctx context.Context, opts ...OptRepo) (*model.User, error)
	List(ctx context.Context, opts ...OptRepo) ([]model.User, error)
	ListRole(ctx context.Context, user_id uint, opts ...OptRepo) ([]model.Role, error)
	Store(ctx context.Context, entity *model.User) error
	StoreUserRole(ctx context.Context, entity *model.UserRole) error
	Update(ctx context.Context, entity *model.User, updates map[string]interface{}) error
}

type user struct {
	db *lib.Database
	*util
}

func NewUserRepository(db *lib.Database) User {
	return &user{db, NewRepoUtil(db)}
}

func (repo *user) Count(ctx context.Context, opts ...OptRepo) (total int64, err error) {
	tx := getDBConnection(ctx, repo.db)
	stmt := tx.Table(repo.GetTableName())
	for _, opt := range opts {
		opt(stmt)
	}
	err = stmt.Count(&total).Error
	if err != nil {
		logger.Error(ctx, "error count "+repo.GetTableName(), map[string]interface{}{
			"error": err.Error(),
			"tags":  []string{"repo", repo.GetTableName(), "count"},
		})
	}
	return
}

func (repo *user) Delete(ctx context.Context, where map[string]interface{}) error {
	tx := getDBConnection(ctx, repo.db)
	err := tx.Where(where).Delete(&model.User{}).Error
	if err != nil {
		logger.Error(ctx, "error delete "+repo.GetTableName(), map[string]interface{}{
			"error": err.Error(),
			"tags":  []string{"repo", "delete", repo.GetTableName()},
		})
	}
	return err
}

func (repo *user) Find(ctx context.Context, opts ...OptRepo) (entity *model.User, err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}
	stmt := tx.Model(&model.User{})
	for _, opt := range opts {
		opt(stmt)
	}
	err = stmt.Find(&entity).Error
	if err != nil {
		logger.Error(ctx, "error find user", map[string]interface{}{
			"error": err.Error(),
			"tags":  []string{"repo", "user"},
		})
		if !errors.Is(gorm.ErrRecordNotFound, err) {
			return
		}
	}
	return entity, nil
}

func (repo *user) GetTableName() string { return "users" }

func (repo *user) List(ctx context.Context, opts ...OptRepo) (entities []model.User, err error) {
	tx := getDBConnection(ctx, repo.db)
	stmt := tx.Model(&model.User{})
	for _, opt := range opts {
		opt(stmt)
	}
	err = stmt.Find(&entities).Error
	if err != nil {
		logger.Error(ctx, "error list user", map[string]interface{}{
			"error": err.Error(),
			"tags":  []string{"repo", "user", "list"},
		})
		if !errors.Is(gorm.ErrRecordNotFound, err) {
			return
		}
	}
	return entities, nil
}

func (repo *user) ListRole(ctx context.Context, user_id uint, opts ...OptRepo) (roles []model.Role, err error) {
	tx := getDBConnection(ctx, repo.db)
	stmt := tx.Select("roles.*").
		Joins("user_roles on user_roles.role_id = roles.role_id").
		Where("user_roles.user_id", user_id)
	for _, opt := range opts {
		opt(stmt)
	}
	err = stmt.Find(&roles).Error
	if err != nil {
		logger.Error(ctx, "user repo: list role "+err.Error(), map[string]interface{}{
			"error": err.Error(),
			"tags":  []string{"repo", repo.GetTableName()},
		})
		if !errors.Is(gorm.ErrRecordNotFound, err) {
			return
		}
	}
	return roles, nil
}

func (repo *user) Store(ctx context.Context, entity *model.User) error {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}
	err := tx.Create(&entity).Error
	if err != nil {
		logger.Error(ctx, "error store user", map[string]interface{}{
			"error": err.Error(),
			"tags":  []string{"repo", "store", "user"},
		})
	}
	return err
}

func (repo *user) StoreUserRole(ctx context.Context, entity *model.UserRole) error {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}
	err := tx.Create(&entity).Error
	if err != nil {
		logger.Error(ctx, "error store user_roles", map[string]interface{}{
			"error": err.Error(),
			"tags":  []string{"repo", "store", "user_roles"},
		})
	}
	return err
}

func (repo *user) Update(ctx context.Context, entity *model.User, updates map[string]interface{}) error {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}
	err := tx.Model(&entity).Updates(updates).Error
	if err != nil {
		logger.Error(ctx, "Error update repo "+repo.GetTableName(), map[string]interface{}{
			"error": err,
			"tags":  []string{"repo", repo.GetTableName(), "repo"},
		})
	}
	return nil
}
