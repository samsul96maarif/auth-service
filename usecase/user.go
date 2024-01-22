package usecase

import (
	"context"

	"github.com/samsul96maarif/auth-service/lib"
	"github.com/samsul96maarif/auth-service/model"
)

func (u *Usecase) FindUser(ctx context.Context, email string) (*model.User, error) {
	user, err := u.repo.FindUser(ctx, map[string]interface{}{
		"email": email,
	}, "")
	if err != nil {
		return nil, err
	}
	if user.Id <= 0 {
		return nil, lib.ErrorNotFound
	}
	return &user, nil
}

func (u *Usecase) FindUserRole(ctx context.Context, user_id int) (*model.UserRole, error) {
	userRole, err := u.repo.FindUserRole(ctx, map[string]interface{}{
		"user_id": user_id,
	}, "")
	if err != nil {
		return nil, err
	}
	if userRole.UserId <= 0 {
		return nil, lib.ErrorNotFound
	}
	return &userRole, nil
}
