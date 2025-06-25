package httpserver

import (
	"encoding/json"
	"net/http"
)

type Server struct {
	walletService WalletService
}

func NewServer(walletService WalletService) *Server {
	return &Server{
		walletService: walletService,
	}
}

func (s *Server) respondWithSuccess(data any, w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (s *Server) respondWithError(httpStatusCode int, message string, errorCode int, w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(httpStatusCode)
	response := ErrorResponse{
		ErrorCode: errorCode,
		Message:   message,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
