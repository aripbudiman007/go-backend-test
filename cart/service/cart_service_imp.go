package service

import (
	"cart/helper.go"
	"cart/model/domain"
	"cart/model/web"
	"cart/repository"
	"context"
	"database/sql"
)

type CartServiceImp struct {
	CartRepository repository.CartRepository
	DB             *sql.DB
}

func (service *CartServiceImp) Save(ctx context.Context, request web.CartCreateRequest) web.CartResponse {
	tx, err := service.DB.Begin()

	helper.PanicIfError(err)

	defer helper.Commit0rRollback(tx)

	cart := domain.Cart{
		ProdukName: request.ProdukName,
	}

	service.CartRepository.Save(ctx, tx, cart)

	return helper.ToCartResponse(cart)
}

func (service *CartServiceImp) Get(ctx context.Context) []web.CartResponse {
	tx, err := service.DB.Begin()

	helper.PanicIfError(err)

	defer helper.Commit0rRollback(tx)

	carts := service.CartRepository.Get(ctx, tx)

	var cartResponse []web.CartResponse

	for _, cart := range carts {
		cartResponse = append(cartResponse, helper.ToCartResponse(cart))
	}

	return cartResponse
}

func (service *CartServiceImp) Delete(ctx context.Context, cartId int) {
	tx, err := service.DB.Begin()

	helper.PanicIfError(err)

	defer helper.Commit0rRollback(tx)

	cart, err := service.CartRepository.FindById(ctx, tx, cartId)

	helper.PanicIfError(err)

	service.CartRepository.Delete(ctx, tx, cart)
}

func (service *CartServiceImp) FindById(ctx context.Context, cartId int) web.CartResponse {
	tx, err := service.DB.Begin()

	helper.PanicIfError(err)

	defer helper.Commit0rRollback(tx)

	cart, err := service.CartRepository.FindById(ctx, tx, cartId)
	helper.PanicIfError(err)

	return helper.ToCartResponse(cart)
}
