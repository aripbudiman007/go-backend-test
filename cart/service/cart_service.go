package service

import (
	"cart/model/web"
	"context"
)

type CartService interface {
	Get(ctx context.Context) []web.CartResponse
	Save(ctx context.Context, request web.CartCreateRequest) web.CartResponse
	FindById(ctx context.Context, cartId int) web.CartResponse
	Delete(ctx context.Context, cartId int)
}
