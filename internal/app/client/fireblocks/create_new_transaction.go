package fireblocks

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *FireblocksClient) CreateNewTransaction(
	ctx context.Context,
	payload CreateNewTransactionRequest,
) (*CreateNewTransactionResponse, error) {
	path := "/v1/transactions" // Can be configurable if needed
	url := c.baseURL + path

	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	token, err := c.signJwt(path, bodyBytes)
	if err != nil {
		return nil, fmt.Errorf("error signing JWT: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-API-KEY", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Check for HTTP errors
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("got error response with status code %d: %s", resp.StatusCode, string(respBody))
	}

	var result CreateNewTransactionResponse
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %w", err)
	}

	return &result, nil
}
