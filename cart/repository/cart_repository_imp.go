package repository

import (
	"cart/helper"
	"cart/model/domain"
	"context"
	"database/sql"
	"errors"
)

type CartRepositoryImp struct {
}

func NewCartRepository() CartRepository {
	return &CartRepositoryImp{}
}

func (repository *CartRepositoryImp) Get(ctx context.Context, tx *sql.Tx) []domain.Cart {
	SQL := "select cart_id, kode_produk, produk_name, kuantiti FROM cart"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()
	var carts []domain.Cart
	for rows.Next() {
		cart := domain.Cart{}
		err := rows.Scan(&cart.CartId, &cart.KodeProduk, &cart.ProdukName, &cart.Kuantiti)
		helper.PanicIfError(err)
		carts = append(carts, cart)
	}

	return carts
}

func (repository *CartRepositoryImp) Save(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart {
	SQL := "insert into cart(kode_produk, produk_name, kuantiti) values (?, ?, ?)"

	result, err := tx.ExecContext(ctx, SQL, cart.KodeProduk, cart.ProdukName, cart.Kuantiti)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()

	helper.PanicIfError(err)

	cart.CartId = int(id)

	return cart
}

func (repository *CartRepositoryImp) FindById(ctx context.Context, tx *sql.Tx, cartId int) (domain.Cart, error) {
	SQL := "select cart_id, kode_produk, produk_name, kuantiti FROM cart where cart_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, cartId)
	helper.PanicIfError(err)
	defer rows.Close()

	cart := domain.Cart{}

	if rows.Next() {
		err := rows.Scan(&cart.CartId, &cart.KodeProduk, &cart.ProdukName, &cart.Kuantiti)
		helper.PanicIfError(err)

		return cart, nil
	} else {
		return cart, errors.New("Product Not Found")
	}
}

func (repository *CartRepositoryImp) Delete(ctx context.Context, tx *sql.Tx, cart domain.Cart) {
	SQL := "delete from cart where cart_id = ?"
	_, err := tx.ExecContext(ctx, SQL, cart.CartId)

	helper.PanicIfError(err)
}
