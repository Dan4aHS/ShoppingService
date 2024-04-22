package entitymodels

type Product struct {
	ID       int     `json:"id" db:"id"`
	UserID   int     `json:"user_id" db:"user_id"`
	Name     string  `json:"name" db:"name"`
	Market   string  `json:"market" db:"market"`
	Price    int     `json:"price" db:"price"`
	Category string  `json:"category" db:"category"`
	Quantity float32 `json:"count" db:"count"`
	Barcode  string  `json:"barcode" db:"barcode"`
}
