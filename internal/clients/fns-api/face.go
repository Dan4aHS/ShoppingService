package fns_api

import (
	"ShoppingExpensesService/internal/models/dbmodels"
)

type IClient interface {
	GetQRInfo(tgID int, qrCodeInfo string) ([]dbmodels.Product, error)
	productsList(tgID int, res FNSResponse) ([]dbmodels.Product, error)
}
