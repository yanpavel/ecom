package products

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yanpavel/api_sandbox/types"
	"github.com/yanpavel/api_sandbox/utils"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleGetProduct).Methods("GET")
}

func (h *Handler) handleGetProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("handleGetProduct called for path:", r.URL.Path)
	ps, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, ps)
}
