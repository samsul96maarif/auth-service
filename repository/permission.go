package repository

import (
	"context"

	"github.com/samsul96maarif/auth-service/lib/logger"
	"github.com/samsul96maarif/auth-service/model"
)

func (repo *Repository) FindPermission(ctx context.Context, opts ...OptRepo) (entity *model.Permission, err error) {
	tx := getDBConnection(ctx, repo.db)
	stmt := tx.Model(&model.Permission{})
	for _, opt := range opts {
		opt(stmt)
	}
	err = stmt.Limit(1).Find(&entity).Error
	if err != nil {
		logger.Error(ctx, "Error find permission", map[string]interface{}{
			"error": err,
			"tags":  []string{"repo", "permission", "repo"},
		})
	}
	return
}

func (repo *Repository) RevokePermission(ctx context.Context, opts ...OptRepo) (entity *model.Permission, err error) {
	tx := getDBConnection(ctx, repo.db)
	stmt := tx.Model(&model.Permission{})
	for _, opt := range opts {
		opt(stmt)
	}
	err = stmt.Limit(1).Find(&entity).Error
	if err != nil {
		logger.Error(ctx, "Error find permission", map[string]interface{}{
			"error": err,
			"tags":  []string{"repo", "permission", "repo"},
		})
	}
	return
}
