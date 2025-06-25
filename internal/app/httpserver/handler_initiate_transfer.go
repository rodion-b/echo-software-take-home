package httpserver

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func (s *Server) InitiateTransferHanlder(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	defer ctx.Done()

	// Check the request method
	if r.Method != http.MethodPost {
		s.respondWithError(http.StatusMethodNotAllowed, "Method Not Allowed", ErrMethodNotAllowed, w, r)
		return
	}

	vars := mux.Vars(r)
	walletID := vars["walletId"]
	// Validate walletID
	if walletID == "" {
		s.respondWithError(http.StatusBadRequest, "walletID is empty", ErrInvalidRequest, w, r)
		return
	}

	// Parse the request body
	var req InitiateTransferRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.respondWithError(http.StatusBadRequest, "Invalid json", ErrInvalidRequest, w, r)
		return
	}
	defer r.Body.Close()

	// Call the service to get the wallet balance
	tx, err := s.walletService.InititateTransfer(
		ctx,
		walletID,
		req.AssetId,
		req.Amount,
		req.SourceAddress,
		req.DestinationAddress,
	)
	if err != nil {
		log.Err(err).Msg("Failed to tansfer")
		s.respondWithError(http.StatusBadRequest, "Failed to tansfer", ErrInternalServerError, w, r)
		return
	}

	// Converting Domain Model to Response Model
	response := InitiateTransferResponse{
		TransactionID:      tx.ID(),
		Status:             tx.Status(),
		AssetId:            tx.AssetID(),
		SourceAddress:      tx.Source(),
		DestinationAddress: tx.Destination(),
		Amount:             tx.Amount(),
	}

	// Respond with success
	s.respondWithSuccess(response, w, r)
}
