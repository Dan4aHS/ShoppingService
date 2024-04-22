package models

import (
	"ShoppingExpensesService/internal/models/dbmodels"
	"ShoppingExpensesService/internal/models/entitymodels"
	shopping "ShoppingExpensesService/pkg"
)

func ProductDBToEntity(prod dbmodels.Product) entitymodels.Product {
	return entitymodels.Product{
		ID:       prod.ID,
		UserID:   prod.UserID,
		Name:     prod.Name,
		Market:   prod.Market,
		Price:    prod.Price,
		Category: prod.Category,
		Quantity: prod.Count,
		Barcode:  prod.Barcode,
	}
}

func ProductEntityToDB(prod entitymodels.Product) dbmodels.Product {
	return dbmodels.Product{
		ID:       prod.ID,
		UserID:   prod.UserID,
		Name:     prod.Name,
		Market:   prod.Market,
		Price:    prod.Price,
		Category: prod.Category,
		Count:    prod.Quantity,
		Barcode:  prod.Barcode,
	}
}

func ProductDBToGRPC(prod dbmodels.Product) *shopping.Purchase {
	return &shopping.Purchase{
		Id:       int32(prod.ID),
		UserId:   int32(prod.UserID),
		Name:     prod.Name,
		Market:   prod.Market,
		Price:    int32(prod.Price),
		Category: prod.Category,
		Quantity: int32(prod.Count),
		Barcode:  prod.Barcode,
	}
}

func ProductsListEntityToDB(list []entitymodels.Product) []dbmodels.Product {
	var result []dbmodels.Product
	for _, prod := range list {
		result = append(result, ProductEntityToDB(prod))
	}
	return result
}

func ProductsListDBToEntity(list []dbmodels.Product) []entitymodels.Product {
	var result []entitymodels.Product
	for _, prod := range list {
		result = append(result, ProductDBToEntity(prod))
	}
	return result
}

func ProductsListDBToGRPC(list []dbmodels.Product) []*shopping.Purchase {
	var result []*shopping.Purchase
	for _, prod := range list {
		result = append(result, ProductDBToGRPC(prod))
	}
	return result
}
