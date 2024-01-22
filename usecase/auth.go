/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2024-01-20 10:18:38
 * @modify date 2024-01-20 10:18:38
 * @desc [description]
 */

package usecase

import (
	"context"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/samsul96maarif/auth-service/config"
	"github.com/samsul96maarif/auth-service/lib"
	"github.com/samsul96maarif/auth-service/lib/logger"
	"github.com/samsul96maarif/auth-service/lib/worker_const"
	"github.com/samsul96maarif/auth-service/model"
	"github.com/samsul96maarif/auth-service/repository"
	"github.com/samsul96maarif/auth-service/request"
	"github.com/samsul96maarif/auth-service/response"
	"golang.org/x/crypto/bcrypt"
)

var (
	DefaultPassword = "password"
	DefaultRoleId   = 4
)

type Auth interface {
	Login(ctx context.Context, r request.Login) (response.Login, error)
	Register(ctx context.Context, r request.Register) (response.User, error)
	ForgotPassword(ctx context.Context, r request.ForgotPassword) error
}

type auth struct {
	util     UsecaseUtil
	userRepo repository.User
}

func NewAuth(util UsecaseUtil, repo repository.User) *auth {
	return &auth{util, repo}
}

func (u *auth) Login(ctx context.Context, req request.Login) (res response.Login, err error) {
	var (
		entity      *model.User
		funcName    = "Login"
		signedToken string
	)
	entity, err = u.userRepo.Find(ctx, repository.OptWithWhere(map[string]interface{}{"email": req.Email}))
	if err != nil {
		return res, lib.CustomInternalServerError("", err.Error(), 0)
	}
	if err != nil {
		return res, lib.CustomInternalServerError("", err.Error(), 0)
	}
	err = bcrypt.CompareHashAndPassword([]byte(entity.Password), []byte(req.Password))
	if err != nil {
		return res, lib.InvalidParameterError("credential", "invalid credential")
	}

	claims := lib.MyClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer: os.Getenv("APP_NAME"),
		},
		UserId: entity.Id,
		Email:  entity.Email,
	}

	token := jwt.NewWithClaims(config.JWT_SIGNING_METHOD, claims)
	if signedToken, err = token.SignedString(config.GetSignatureKey()); err != nil {
		return res, lib.CustomInternalServerError("token", err.Error(), 0)
	}
	res.Token = signedToken
	go func() {
		if err = u.util.DispatchWorker(worker_const.KeyWorkerCache, request.Cache{
			Key:     signedToken,
			Data:    claims,
			Expires: 24 * time.Hour,
		}); err != nil {
			logger.Error(ctx, funcName+" dispatch worker error: "+err.Error(), nil)
		}
	}()
	return res, nil
}

func (u *auth) Register(ctx context.Context, r request.Register) (res response.User, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.MinCost)
	if err != nil {
		return res, lib.InvalidParameterError("password", err.Error())
	}
	entity := model.User{
		Name:     r.Name,
		Email:    r.Email,
		Password: string(hash),
	}
	err = u.userRepo.Transaction(ctx, func(ctx context.Context) error {
		if err = u.userRepo.Store(ctx, &entity); err != nil {
			return err
		}
		userRoleEntity := model.UserRole{
			UserId: entity.Id,
			RoleId: uint(DefaultRoleId),
		}
		if err = u.userRepo.StoreUserRole(ctx, &userRoleEntity); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}

	res = response.NewUser(entity, nil)
	return
}

func (u *auth) ForgotPassword(ctx context.Context, r request.ForgotPassword) error {
	panic("not implemented") // TODO: Implement
}
