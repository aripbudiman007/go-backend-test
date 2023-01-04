package repository

import (
	"cart/model/domain"
	"context"
	"database/sql"
)

type CartRepository interface {
	Get(ctx context.Context, tx *sql.Tx) []domain.Cart
	Save(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart
	FindById(ctx context.Context, tx *sql.Tx, cartId int) (domain.Cart, error)
	Delete(ctx context.Context, tx *sql.Tx, cart domain.Cart)
}
