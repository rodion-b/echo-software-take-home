package fireblocks

import (
	"crypto/rsa"
	"crypto/sha256"
	"echo-software-take-home/internal/app/config"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// FireblocksClient represents a client for interacting with the Fireblocks API
type FireblocksClient struct {
	baseURL    string
	apiKey     string
	privateKey *rsa.PrivateKey
	httpClient *http.Client
}

// NewFireblocksClient creates a new Fireblocks client instance
func NewFireblocksClient(config config.Config) (*FireblocksClient, error) {
	privateKeyBytes, err := os.ReadFile(config.SECRET_KEY_PATH)
	if err != nil {
		return nil, fmt.Errorf("error reading private key from %s: %w", config.SECRET_KEY_PATH, err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing RSA private key: %w", err)
	}

	return &FireblocksClient{
		baseURL:    config.BASE_URL,
		apiKey:     config.API_KEY,
		privateKey: privateKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second, // TO-DO: Make configurable
		},
	}, nil
}

func (a *FireblocksClient) signJwt(path string, bodyBytes []byte) (string, error) {
	nonce := uuid.New().String()
	now := time.Now().Unix()
	expiration := now + 55 // Consider making this configurable

	h := sha256.New()
	h.Write(bodyBytes)
	hashed := h.Sum(nil)

	claims := jwt.MapClaims{
		"uri":      path,
		"nonce":    nonce,
		"iat":      now,
		"exp":      expiration,
		"sub":      a.apiKey,
		"bodyHash": hex.EncodeToString(hashed),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(a.privateKey)
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return tokenString, nil
}
