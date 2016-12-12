package goswrve

// Client represent an API client for the Swrve service.
type Client struct {
	error *APIError
}

// GetError retrieves the last client error.
func (c *Client) GetError() *APIError {
	return c.error
}
