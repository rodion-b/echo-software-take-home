package httpserver

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func (s *Server) CreateWalletHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	defer ctx.Done()

	// Check if the request method is POST
	if r.Method != http.MethodPost {
		s.respondWithError(http.StatusMethodNotAllowed, "Method Not Allowed", ErrMethodNotAllowed, w, r)
		return
	}

	// Parse the request body
	var req CreateWalletRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.respondWithError(http.StatusBadRequest, "Invalid json", ErrInvalidRequest, w, r)
		return
	}
	defer r.Body.Close()

	// Validate request
	// TO-DO: Use a validation library
	if req.Name == "" {
		s.respondWithError(http.StatusBadRequest, "Wallet name is required", ErrInvalidRequest, w, r)
		return
	}

	// Create the wallet using the wallet service
	wallet, err := s.walletService.CreateWallet(ctx, req.Name)
	if err != nil {
		log.Err(err).Msg("Failed to create wallet")
		s.respondWithError(http.StatusBadRequest, "Failed to create wallet", ErrNewWalletCreation, w, r)
		return
	}

	// Converting Domain Model to Response Model
	response := CreateWalletResponse{
		ID:      wallet.ID(),
		Name:    wallet.Name(),
		VaultID: wallet.VaultID(),
	}

	// Respond with success
	s.respondWithSuccess(response, w, r)
}
