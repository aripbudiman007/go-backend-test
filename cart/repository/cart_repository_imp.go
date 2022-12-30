package repository

import (
	"cart/helper.go"
	"cart/model/domain"
	"context"
	"database/sql"
	"errors"
)

type CartRepositoryImp struct {
}

func (repository *CartRepositoryImp) Get(ctx context.Context, tx *sql.Tx, cart domain.Cart) []domain.Cart {
	SQL := "select cart_id, produk_name, kuatiti from cart"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var carts []domain.Cart
	for rows.Next() {
		cart := domain.Cart{}
		err := rows.Scan(cart.CartId, &cart.ProdukName)
		helper.PanicIfError((err))
		carts = append(carts, cart)
	}

	return carts
}

func (repository *CartRepositoryImp) Save(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart {
	SQL := "insert into cart(produk_name) value(?)"
	result, err := tx.ExecContext(ctx, SQL, cart.ProdukName)

	helper.PanicIfError(err)

	id, err := result.LastInsertId()

	helper.PanicIfError(err)

	cart.CartId = int(id)

	return cart
}

func (repository *CartRepositoryImp) Delete(ctx context.Context, tx *sql.Tx, cart domain.Cart) {
	SQL := "delete from cart where cart_id = ?"
	_, err := tx.ExecContext(ctx, SQL, cart.CartId)

	helper.PanicIfError(err)
}

func (repository *CartRepositoryImp) FindById(ctx context.Context, tx *sql.Tx, cartId int) (domain.Cart, error) {
	SQL := "select cart_id, produk_name from cart where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, cartId)
	helper.PanicIfError(err)

	cart := domain.Cart{}
	if rows.Next() {
		err := rows.Scan(&cart.CartId, &cart.ProdukName)
		helper.PanicIfError(err)

		return cart, nil
	} else {
		return cart, errors.New("Product Not Found")
	}
}
