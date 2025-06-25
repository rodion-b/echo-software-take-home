package httpserver

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func (s *Server) GetDepositAddressHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	defer ctx.Done()

	vars := mux.Vars(r)
	walletID := vars["walletId"]
	assetID := vars["assetId"]

	// Check the request method
	if r.Method != http.MethodGet {
		s.respondWithError(http.StatusMethodNotAllowed, "Method Not Allowed", ErrMethodNotAllowed, w, r)
		return
	}

	// Validate request
	if walletID == "" {
		s.respondWithError(http.StatusBadRequest, "walletID is empty", ErrInvalidRequest, w, r)
		return
	}
	if assetID == "" {
		s.respondWithError(http.StatusBadRequest, "assetID is empty", ErrInvalidRequest, w, r)
		return
	}

	// Call the wallet service
	addresses, err := s.walletService.GetDepositAddress(ctx, walletID, assetID)
	if err != nil {
		log.Err(err).Msg("Failed to get deposit address")
		s.respondWithError(http.StatusBadRequest, "Failed to get deposit address", ErrInternalServerError, w, r)
		return
	}

	addressIds := make([]string, 0, len(addresses))
	for _, addr := range addresses {
		addressIds = append(addressIds, addr.Address())
	}

	// Converting Domain Model to Response Model
	response := GetDepositAddressResponse{
		WalletID:  walletID,
		AssetID:   assetID,
		Addresses: addressIds,
	}

	// Respond with success
	s.respondWithSuccess(response, w, r)
}
