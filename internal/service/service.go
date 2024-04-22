package service

import (
	"ShoppingExpensesService/internal/models"
	"ShoppingExpensesService/internal/models/dbmodels"
	"ShoppingExpensesService/internal/repo"
	"context"
	"log"
)

type Service struct {
	Repo repo.IProductRepository `json:"repo"`
}

func NewService(repo repo.IProductRepository) *Service {
	err := repo.CreateTable()
	if err != nil {
		log.Fatal(err)
	}
	return &Service{Repo: repo}
}

func (s Service) AddProducts(ctx context.Context, products []dbmodels.Product) (int, error) {
	count, err := s.Repo.AddProducts(ctx, models.ProductsListDBToEntity(products))
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s Service) ListProducts(ctx context.Context, tgID int) ([]dbmodels.Product, error) {
	products, err := s.Repo.ListProducts(ctx, tgID)
	if err != nil {
		return nil, err
	}
	return models.ProductsListEntityToDB(products), nil
}
