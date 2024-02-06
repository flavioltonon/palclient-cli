package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
}

func NewClient(baseURL *url.URL, httpClient *http.Client) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: httpClient,
	}
}

type ExecuteWebhookRequest struct {
	Content string `json:"content"`
}

type ExecuteWebhookResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e ExecuteWebhookResponseError) Error() string {
	return fmt.Sprintf("%s (code: %d)", e.Message, e.Code)
}

func (c *Client) ExecuteWebhook(message string) error {
	request := ExecuteWebhookRequest{
		Content: message,
	}

	b, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("marshalling body: %w", err)
	}

	response, err := http.Post(c.baseURL.String(), "application/json", bytes.NewReader(b))
	if err != nil {
		return fmt.Errorf("making request: %w", err)
	}

	if response.StatusCode != http.StatusNoContent {
		var responseErr ExecuteWebhookResponseError

		if err := json.NewDecoder(response.Body).Decode(&responseErr); err != nil {
			return fmt.Errorf("decoding error response: %w", err)
		}

		return responseErr
	}

	return nil
}
