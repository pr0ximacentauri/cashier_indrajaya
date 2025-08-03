package handlers

import(
	"cashier_indrajaya/usecases"
	"cashier_indrajaya/domain/entities"
	"encoding/json"
	"net/http"
)

type SaleHandler struct {
	service *usecases.SaleService
}

func NewSaleHandler(service *usecases.SaleService) *SaleHandler {
	return &SaleHandler{service: service}
}

func (h *SaleHandler) CreateSale(w http.ResponseWriter, r *http.Request) {
	var sale entities.Sale
	if err := json.NewDecoder(r.Body).Decode(&sale); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateSale(&sale); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sale)
}

func (h *SaleHandler) GetSales(w http.ResponseWriter, r *http.Request) {
	sales, err := h.service.GetAllSales()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sales)
}