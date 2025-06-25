package fireblocks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *FireblocksClient) GetAssetBalanceForVaultAccountId(
	ctx context.Context,
	vaultAccountId string,
	assetId string,
) (*GetAssetBalanceForVaultResponse, error) {
	path := "/v1/vault/accounts/" // Base path for vault accounts
	fullPath := path + fmt.Sprintf("%s/%s", vaultAccountId, assetId)

	url := c.baseURL + fullPath

	token, err := c.signJwt(fullPath, nil)
	if err != nil {
		return nil, fmt.Errorf("error signing JWT: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %w", err)
	}

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

	var result GetAssetBalanceForVaultResponse
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %w", err)
	}

	return &result, nil
}
