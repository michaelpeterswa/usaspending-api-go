package usaspending

import (
	"net/http"
	"time"
)

const (
	USASpendingBaseURL = "https://api.usaspending.gov/api/v2"
)

type USASpendingClient struct {
	httpClient *http.Client
}

func NewUSASpendingClient(timeout time.Duration) *USASpendingClient {
	return &USASpendingClient{
		httpClient: &http.Client{Timeout: timeout},
	}
}

func NewCustomUSASpendingClient(client *http.Client) *USASpendingClient {
	return &USASpendingClient{
		httpClient: client,
	}
}
