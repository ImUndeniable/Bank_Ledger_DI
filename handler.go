package main

import (
	"encoding/json"
	"net/http"
)

type TransferRequest struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

// Dependency Injected - conn to service layer
type HandlerService struct {
	service TransferService
}

// Injector

func NewHandlerService(h TransferService) *HandlerService {
	return &HandlerService{
		service: h,
	}
}

func (h *HandlerService) HandleTransfer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	var req TransferRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := h.service.SendMoney(req.From, req.To, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Transfer Successfull!\n"))
}
