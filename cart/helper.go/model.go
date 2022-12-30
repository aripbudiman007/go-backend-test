package helper

import (
	"cart/model/domain"
	"cart/model/web"
)

func ToCartResponse(cart domain.Cart) web.CartResponse {
	return web.CartResponse{
		KdProduk:   cart.CartId,
		ProdukName: cart.ProdukName,
		Kuantiti:   cart.Kuantiti,
	}
}
