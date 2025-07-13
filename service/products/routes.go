package products

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/yanpavel/api_sandbox/service/auth"
	"github.com/yanpavel/api_sandbox/types"
	"github.com/yanpavel/api_sandbox/utils"
)

type Handler struct {
	store     types.ProductStore
	userStore types.UserStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", auth.WithJWTAuth(h.handleGetProducts, h.userStore)).Methods("GET")
	router.HandleFunc("/products", auth.WithJWTAuth(h.handleCreateProduct, h.userStore)).Methods("POST")
	router.HandleFunc("/products/{productID}", auth.WithJWTAuth(h.handleGetProductById, h.userStore)).Methods("GET")
	router.HandleFunc("/products", auth.WithJWTAuth(h.handleUpdateProducts, h.userStore)).Methods("PUT")
}

func (h *Handler) handleUpdateProducts(w http.ResponseWriter, r *http.Request) {
	var payload types.UpdateProductPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("invalid type passed to validator: %v", err))
			return
		}

		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	err := h.store.UpdateProduct(types.Product{
		ID:          payload.ID,
		Name:        payload.Name,
		Price:       *payload.Price,
		Description: *payload.Description,
		Image:       *payload.Image,
		Quantity:    *payload.Quantity,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, payload.ID)
}

func (h *Handler) handleGetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["productID"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing product ID"))
		return
	}

	productID, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid format of productID"))
		return
	}

	pt, err := h.store.GetProductByID(productID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("product is not existed"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, pt)

}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateProductPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	//validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("invalid type passed to validator: %v", err))
			return
		}

		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	err := h.store.CreateProducts(types.Product{
		Name:        payload.Name,
		Description: payload.Description,
		Image:       payload.Image,
		Price:       payload.Price,
		Quantity:    payload.Quantity,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	ps, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, ps)
}
