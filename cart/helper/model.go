package helper

import (
	"cart/model/domain"
	"cart/model/web"
)

func ToCartResponse(cart domain.Cart) web.CartResponse {
	return web.CartResponse{
		CartId:     cart.CartId,
		KodeProduk: cart.KodeProduk,
		ProdukName: cart.ProdukName,
		Kuantiti:   cart.Kuantiti,
	}
}
