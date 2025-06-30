package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/yanpavel/api_sandbox/types"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if the user payload is invalid",
		func(t *testing.T) {
			payload := types.RegisterUserPayload{
				FirstName: "user3",
				LastName:  "123 ",
				Email:     "asd",
				Password:  "asd4",
			}
			marshalled, _ := json.Marshal(payload)

			req, err := http.NewRequest(http.MethodPost,
				"/register", bytes.NewBuffer(marshalled))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router := mux.NewRouter()

			router.HandleFunc("/register", handler.handleRegister)
			router.ServeHTTP(rr, req)
			fmt.Print(payload)
			if rr.Code != http.StatusBadRequest {
				t.Errorf("expected status code %d, got %d",
					http.StatusBadRequest, rr.Code)
			}
		})

	t.Run("should correctly register the user", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "user2",
			LastName:  "123 ",
			Email:     "asd2@gmail.com",
			Password:  "asd4",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost,
			"/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)
		fmt.Print(payload)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d",
				http.StatusCreated, rr.Code)
		}
	})
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(types.User) error {
	return nil
}
