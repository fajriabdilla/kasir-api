package handlers

import (
	"encoding/json"
	"kasir-api/models"
	"kasir-api/services"
	"log"
	"net/http"
)

type TransactionHandler struct {
	service *services.TransactionService
}

func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	log.Println("[HANDLER] Checkout called")

	var req models.CheckoutRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	// log.Println(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// semua item ke service transaction
	transaction, err := h.service.Checkout(req.Items)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)

}

func (h *TransactionHandler) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Checkout(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
