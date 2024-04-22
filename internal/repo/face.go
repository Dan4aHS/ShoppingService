package repo

import (
	"ShoppingExpensesService/internal/models/entitymodels"
	"context"
)

type IProductRepository interface {
	CreateTable() error
	AddProducts(ctx context.Context, products []entitymodels.Product) (int, error)
	AddProduct(ctx context.Context, prod entitymodels.Product) error
	ListProducts(ctx context.Context, tgID int) ([]entitymodels.Product, error)
	Rollback(prod entitymodels.Product) error
	Commit(prod entitymodels.Product) error
}
