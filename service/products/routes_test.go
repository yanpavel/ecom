package products

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/gorilla/mux"
	"github.com/yanpavel/api_sandbox/types"
)

type mockProductStore struct {
	mx sync.RWMutex
}

func (m *mockProductStore) UpdateProduct(product types.Product) error {
	m.mx.Lock()
	defer m.mx.Unlock()

	product.Quantity++
	return nil
}

func (m *mockProductStore) CreateProducts(types.Product) error {
	return nil
}
func (m *mockProductStore) GetProductsByID([]int) ([]types.Product, error) {
	return nil, nil
}

func (m *mockProductStore) GetProductByID(int) (*types.Product, error) {
	return nil, nil
}

func (m *mockProductStore) GetProducts() ([]*types.Product, error) {
	return nil, nil
}

func TestUpdateProducts(t *testing.T) {
	store := &mockProductStore{}
	handler := NewHandler(store)
	t.Run("Проверка метода обновления продуктов", func(t *testing.T) {
		payload := types.Product{
			ID:       1,
			Quantity: 2,
		}
		marshaled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPut, "/products", bytes.NewBuffer(marshaled))
		if err != nil {
			t.Errorf("Ошибка при отправке запроса %v", err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/products", handler.handleUpdateProducts).Methods("PUT")
		router.ServeHTTP(rr, req)
		fmt.Println(payload)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected status code %v, got %v", http.StatusOK, rr.Code)
		}
		t.Logf("Status code %v", rr.Code)
	})
	t.Run("Проверка метода получения продуктов", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/products", nil)
		if err != nil {
			t.Errorf("Ошибка при отправке запроса %v", err)
			return
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/products", handler.handleGetProducts).Methods("GET")
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected status code %v, got %v", http.StatusOK, rr.Code)
		}
		t.Logf("Status code %v", rr.Code)
	})
	t.Run("Проверка создания продукта", func(t *testing.T) {
		payload := types.CreateProductPayload{
			Name:        "Product1",
			Description: "superpeuper",
			Price:       100.99,
			Quantity:    50,
		}
		marshaled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(marshaled))
		if err != nil {
			t.Errorf("Ошибка при отправке запроса %v", err)
			return
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/products", handler.handleCreateProduct).Methods("POST")
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("Expected status code %v, got %v", http.StatusCreated, rr.Code)
		}
	})
	t.Run("Проверка асинхронной обработки", func(t *testing.T) {

		product := types.Product{
			ID:       3,
			Quantity: 1,
		}
		quantity := &product.Quantity
		result := 101
		const numGoroutines = 101

		var wg sync.WaitGroup
		wg.Add(numGoroutines)

		for range numGoroutines {
			go func() {
				*quantity += 1
				store.UpdateProduct(product)

			}()
			wg.Done()
		}

		wg.Wait()
		if *quantity != result {
			t.Errorf("Expected:%v got %d", result, *quantity)
		}
	})
	t.Run("Проверка получения продукта", func(t *testing.T) {

	})
}
