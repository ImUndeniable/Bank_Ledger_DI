package main

type TransferRequest struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount string `json:"amount"`
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
