package products

import (
	"database/sql"
	"fmt"

	"github.com/yanpavel/api_sandbox/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateProducts(p *types.Product) error {
	_, err := s.db.Exec("INSERT INTO products (name, description, image, price, quantity, createdAt) VALUES (?, ?, ?, ?, ?, ?)",
		p.Name, p.Description, p.Image, p.Price, p.Quantity, p.CreatedAt)

	if err != nil {
		return fmt.Errorf("Product isn't created : %v", err)
	}
	return nil
}

func (s *Store) GetProducts() ([]*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]*types.Product, 0)
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}
	return products, nil
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	products := new(types.Product)

	err := rows.Scan(
		&products.ID,
		&products.Name,
		&products.Price,
		&products.Description,
		&products.Image,
		&products.CreatedAt,
		&products.Quantity,
	)

	if err != nil {
		return nil, err
	}

	return products, nil
}
