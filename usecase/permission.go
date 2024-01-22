package usecase

import (
	"context"
	"fmt"

	"github.com/samsul96maarif/auth-service/lib"
	"github.com/samsul96maarif/auth-service/model"
	"github.com/samsul96maarif/auth-service/repository"
)

const (
	cachePermission = "*"
)

func (u *Usecase) Hello(ctx context.Context) (string, error) {
	var skipPermission bool
	url := ctx.Value("url")
	httpMethod := ctx.Value("http_method")
	email := ctx.Value("email")
	user, err := u.repo.FindUser(ctx, map[string]interface{}{
		"email": email,
	}, "")
	if err != nil {
		return "", err
	}
	userRole, err := u.repo.FindUserRole(ctx, map[string]interface{}{
		"user_id": user.Id,
	}, "")
	if err != nil {
		return "", err
	}
	if userRole.RoleId == lib.ROLE_SUPER_ADMIN_ID {
		skipPermission = true
	}
	if skipPermission {
		return "sukses", nil
	}
	permission, err := u.repo.FindPermission(ctx, repository.OptWithWhere(map[string]interface{}{
		"url":         url,
		"http_method": httpMethod,
	}))
	if err != nil {
		return "", lib.ErrorForbidden
	}
	// ini bisa di cache
	rolePermission, err := u.repo.FindRolePermission(ctx, repository.OptWithWhere(map[string]interface{}{
		"role_id":         userRole.RoleId,
		"permission_slug": permission.Slug,
	}))
	if err != nil {
		fmt.Println("Error", err)
		return "", err
	}
	if rolePermission.RoleId <= 0 {
		return "", lib.ErrorForbidden
	}
	fmt.Println("here", rolePermission)
	fmt.Sprintf("%+v \n", rolePermission)
	return "sukses", nil
}

func (u *Usecase) FindRolePermission(ctx context.Context, role_id uint, permission_slug string) (*model.RolePermission, error) {
	entity, err := u.repo.FindRolePermission(ctx, repository.OptWithWhere(map[string]interface{}{
		"role_id":         role_id,
		"permission_slug": permission_slug,
	}))
	if err != nil {
		return nil, err
	}
	if entity.RoleId <= 0 {
		return nil, lib.ErrorForbidden
	}
	return entity, nil
}

func (u *Usecase) FindPermission(ctx context.Context, route_slag string) (*model.Permission, error) {
	entity, err := u.repo.FindPermission(ctx, repository.OptWithWhere(map[string]interface{}{
		"slag": route_slag,
	}))
	if err != nil {
		return nil, err
	}
	if entity.Slug == "" {
		return nil, lib.ErrorNotFound
	}
	return entity, nil
}

func (u *Usecase) AdddPermission(ctx context.Context, role_id uint, permission_slug string) error {
	entity := model.RolePermission{RoleId: role_id, PermissionSlug: permission_slug}
	err := u.repo.CreateRolePermission(ctx, &entity)
	if err != nil {
		return err
	}
	return nil
}

func (u *Usecase) RevokePermission(ctx context.Context, role_id uint, permission_slug string) error {
	err := u.repo.DeleteRolePermission(ctx, role_id, permission_slug)
	if err != nil {
		return err
	}
	return nil
}
