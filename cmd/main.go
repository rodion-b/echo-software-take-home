package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"echo-software-take-home/internal/app/client/fireblocks"
	"echo-software-take-home/internal/app/config"
	"echo-software-take-home/internal/app/httpserver"
	"echo-software-take-home/internal/app/service"
	"echo-software-take-home/internal/app/storage/pg"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := run(); err != nil {
		log.Err(err).Msg("Error in run")
	}
	os.Exit(0)
}

func run() error {
	// Load environment variables from .env file for testing purposes only
	// App should be dockerised for production use - to be removed in production
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("Error loading .env file: %v", err)
	}

	// Initialize configuration
	config := config.NewConfig()

	// Create a new Fireblocks client
	fireClient, err := fireblocks.NewFireblocksClient(config)
	if err != nil {
		return fmt.Errorf("Error creating Fireblocks client: %v", err)
	}

	// DEBUG ONLY: Get all vault accounts from Fireblocks
	acccounts, err := fireClient.GetAccountsPaged(context.Background())
	if err != nil {
		return fmt.Errorf("Error getting accounts: %v", err)
	}
	fmt.Printf("%+v\n", acccounts)

	// Init databse connection
	db, err := pg.NewPgStorage(config)

	// Create a Wallet Service here
	walletService := service.NewWalletService(fireClient, db)

	//Create a new Server wrapper
	server := httpserver.NewServer(walletService)

	r := mux.NewRouter()
	r.HandleFunc("/wallets", server.CreateWalletHandler)
	r.HandleFunc("/wallets/{walletId}/assets/{assetId}/balance", server.GetWalletBalanceHandler)
	r.HandleFunc("/wallets/{walletId}/assets/{assetId}/address", server.GetDepositAddressHandler)
	r.HandleFunc("/wallets/{walletId}/transactions", server.InitiateTransferHanlder)

	// Create server with timeouts
	srv := &http.Server{
		Addr:         config.HOST,
		Handler:      r,
		ReadTimeout:  15 * time.Second, // TO-DO: Make configurable
		WriteTimeout: 15 * time.Second, // TO-DO: Make configurable
		IdleTimeout:  60 * time.Second, // TO-DO: Make configurable
	}

	// listen to OS signals and gracefully shutdown server
	stopped := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Err(err).Msg("HTTP Server Shutdown Error")
		}
		close(stopped)
	}()

	log.Info().Msg(fmt.Sprintf("Starting HTTP server on %s", config.HOST))

	// start HTTP server
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		return fmt.Errorf("HTTP server ListenAndServe Error: %v", err)
	}

	<-stopped
	log.Info().Msg("Server gracefully stopped")

	return nil
}
