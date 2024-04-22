package service

import (
	"ShoppingExpensesService/internal/models/dbmodels"
	"context"
)

type IProductService interface {
	AddProducts(ctx context.Context, products []dbmodels.Product) (int, error)
	ListProducts(ctx context.Context, tgID int) ([]dbmodels.Product, error)
}
