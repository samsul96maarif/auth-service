/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2024-01-18 17:42:01
 * @modify date 2024-01-18 17:42:01
 * @desc [description]
 */
package usecase

import (
	"context"

	"github.com/samsul96maarif/auth-service/lib"
	"github.com/samsul96maarif/auth-service/repository"
	"github.com/samsul96maarif/auth-service/request"
	"github.com/samsul96maarif/auth-service/response"
)

type BaseUsecase interface {
	Delete(ctx context.Context, id uint) error
	Find(ctx context.Context, req request.FindRequest) (interface{}, error)
	Store(ctx context.Context, req request.StoreRequest) (interface{}, error)
}

type ListUseCase interface {
	List(ctx context.Context, req request.ListReq) (interface{}, error)
}

type PaginationUseCase interface {
	Get(ctx context.Context, req request.PaginateReq) (response.PaginateResponse, error)
}

type UpdateUseCase interface {
	Update(ctx context.Context, req request.UpdateRequest) (interface{}, error)
}

type Usecase struct {
	repo *repository.Repository
}

func NewUsecase(db *lib.Database) Usecase {
	repo := repository.NewRepository(db)
	return Usecase{repo: &repo}
}
