package postgres

import (
	"ShoppingExpensesService/internal/models/entitymodels"
	"context"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type Repository struct {
	db *sqlx.DB
	ts map[int]*sqlx.Tx
}

func NewRepository(db *sqlx.DB) *Repository {
	slog.Info("Created new repository")
	return &Repository{
		db: db,
		ts: make(map[int]*sqlx.Tx),
	}
}

func (r Repository) CreateTable() error {
	q := `CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY ,
    user_id INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    market VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL,
    category VARCHAR(255) NOT NULL,
    quantity INTEGER NOT NULL,
    barcode VARCHAR(255) NOT NULL
)`
	_, err := r.db.Exec(q)
	if err != nil {
		slog.Error("Creating Table failed", err)
		return err
	}
	slog.Info("Table created")
	return nil
}

func (r Repository) AddProducts(ctx context.Context, products []entitymodels.Product) (int, error) {
	if len(products) == 0 {
		return 0, nil
	}
	var err error
	if _, ok := r.ts[products[0].ID]; ok {
		r.ts[products[0].ID], err = r.db.BeginTxx(ctx, nil)
		if err != nil {
			return 0, err
		}
		slog.InfoContext(ctx, "Started transaction")
	}
	for _, product := range products {
		err = r.AddProduct(ctx, product)
		if err != nil {
			slog.ErrorContext(ctx, "Error adding product", err)
			rbErr := r.Rollback(product)
			if rbErr != nil {
				return 0, rbErr
			}
			return 0, err
		}
	}
	err = r.Commit(products[0])
	if err != nil {
		return 0, err
	}
	slog.InfoContext(ctx, "Committed transaction")
	return len(products), nil
}

func (r Repository) AddProduct(ctx context.Context, prod entitymodels.Product) error {
	q := `
	INSERT INTO products
		(user_id, name, market, price, category, quantity, barcode)
	VALUES 
	    ($1, $2, $3, $4, $5, $6, $7)
`
	var err error
	if _, ok := r.ts[prod.UserID]; !ok {
		r.ts[prod.UserID], err = r.db.BeginTxx(ctx, nil)
		if err != nil {
			return err
		}
	}
	_, err = r.ts[prod.UserID].ExecContext(
		ctx,
		q,
		prod.UserID,
		prod.Name,
		prod.Market,
		prod.Price,
		prod.Category,
		prod.Quantity,
		prod.Barcode,
	)
	if err != nil {
		slog.ErrorContext(ctx, "Error adding product", err)
		return err
	}
	slog.InfoContext(ctx, "Added product", prod.Name)
	return nil
}

func (r Repository) ListProducts(ctx context.Context, tgID int) ([]entitymodels.Product, error) {
	var products []entitymodels.Product
	q := `
		SELECT 
		    id, user_id, name, market, price, category, quantity, barcode
		FROM 
		    products
		WHERE 
		    user_id = $1
`
	err := r.db.SelectContext(ctx, &products, q, tgID)
	if err != nil {
		slog.ErrorContext(ctx, "Error listing products", err)
		return nil, err
	}
	slog.InfoContext(ctx, "Success listing products from db")
	return products, nil
}

func (r Repository) Rollback(prod entitymodels.Product) error {
	if t, ok := r.ts[prod.UserID]; ok {
		if err := t.Rollback(); err != nil {
			return err
		}
		delete(r.ts, prod.UserID)
	}
	return nil
}

func (r Repository) Commit(prod entitymodels.Product) error {
	if t, ok := r.ts[prod.UserID]; ok {
		if err := t.Commit(); err != nil {
			return err
		}
		delete(r.ts, prod.UserID)
	}
	return nil
}
