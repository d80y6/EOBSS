package netbox

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// NetBoxClient provides integration with NetBox IPAM/DCIM
type NetBoxClient struct {
	BaseURL    string
	APIToken   string
	httpClient *http.Client
}

func NewNetBoxClient(baseURL, token string) *NetBoxClient {
	return &NetBoxClient{
		BaseURL:  baseURL,
		APIToken: token,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetDeviceBySerialNumber fetches device details from NetBox
func (c *NetBoxClient) GetDeviceBySerialNumber(ctx context.Context, serial string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/dcim/devices/?serial=%s", c.BaseURL, serial)
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", c.APIToken))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("netbox error: %d", resp.StatusCode)
	}

	// In production, we'd unmarshal to a typed struct
	return map[string]interface{}{"status": "success"}, nil
}
