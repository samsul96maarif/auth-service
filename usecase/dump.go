package usecase

import (
	"context"

	"github.com/samsul96maarif/auth-service/lib"
	"github.com/samsul96maarif/auth-service/model"
	"golang.org/x/crypto/bcrypt"
)

func (u *Usecase) CreateSuperAdmin(ctx context.Context) (err error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte("localhost"), bcrypt.MinCost)
	user := model.User{
		Email:    "super-admin@maarif.id",
		Password: string(hash),
	}
	entities, err := u.repo.GetUser(ctx, map[string]interface{}{}, "")
	if len(entities) <= 0 {
		err = u.repo.Transaction(ctx, func(ctx context.Context) error {
			if err = u.repo.CreateUser(ctx, &user); err != nil {
				return err
			}
			err = u.repo.CreateUserRole(ctx, &model.UserRole{
				UserId: user.Id,
				RoleId: lib.ROLE_SUPER_ADMIN_ID,
			})

			return err
		})
	}
	if err != nil {
		err = lib.CustomInternalServerError("", err.Error(), 0)
	}
	return err
}
