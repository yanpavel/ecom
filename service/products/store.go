package products

import (
	"database/sql"
	"fmt"
	"strings"
	"sync"

	"github.com/yanpavel/api_sandbox/types"
)

type Store struct {
	db *sql.DB
	mx sync.RWMutex
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateProducts(p types.Product) error {
	_, err := s.db.Exec("INSERT INTO products (name, description, image, price, quantity) VALUES (?, ?, ?, ?, ?)",
		p.Name, p.Description, p.Image, p.Price, p.Quantity)

	if err != nil {
		return fmt.Errorf("product isn't created : %v", err)
	}
	return nil
}

func (s *Store) GetProductByID(ID int) (*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products WHERE Id=?", ID)
	if err != nil {
		return nil, err
	}
	product, err := scanRowsIntoProduct(rows)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *Store) GetProductsByID(productIDs []int) ([]types.Product, error) {
	s.mx.RLock()
	defer s.mx.RUnlock()

	placeholders := strings.Repeat(",?", len(productIDs)-1)
	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (?%s)", placeholders)

	args := make([]interface{}, len(productIDs))
	for i, v := range productIDs {
		args[i] = v
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	products := []types.Product{}
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func (s *Store) GetProducts() ([]*types.Product, error) {
	s.mx.RLock()
	defer s.mx.RUnlock()

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

func (s *Store) UpdateProduct(product types.Product) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	_, err := s.db.Exec("UPDATE products SET name = ?, price = ?, image =?, description = ?, quantity = ? WHERE id = ?",
		product.Name, product.Price, product.Image, product.Description, product.Quantity, product.ID)
	if err != nil {
		return err
	}
	return nil
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
