package httpserver

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)
// Example request: localhost:8080/wallets/10/assets/BTC_TEST/balance
func (s *Server) GetWalletBalanceHandler(w http.ResponseWriter, r *http.Request) {
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

	// Validate walletID
	if walletID == "" {
		s.respondWithError(http.StatusBadRequest, "walletID is empty", ErrInvalidRequest, w, r)
		return
	}
	// Validate assetID
	if assetID == "" {
		s.respondWithError(http.StatusBadRequest, "assetID is empty", ErrInvalidRequest, w, r)
		return
	}

	// Call the service to get the wallet balance
	asset, err := s.walletService.GetWalletBalance(ctx, walletID, assetID)
	if err != nil {
		log.Err(err).Msg("Failed to get wallet balance")
		s.respondWithError(http.StatusBadRequest, "Failed to get wallet balance", ErrInternalServerError, w, r)
		return
	}

	// Converting Domain Model to Response Model
	response := GetWalletBalanceResponse{
		WalletID:         walletID,
		AssetID:          asset.ID(),
		BalanceTotal:     asset.BalanceTotal(),
		BalanceAvailable: asset.BalanceAvailable(),
	}

	// Respond with success
	s.respondWithSuccess(response, w, r)
}
