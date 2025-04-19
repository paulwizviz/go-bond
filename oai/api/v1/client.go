package oaiapi

import (
	"net/http"
	"time"
)

// Client is an interface for the OpenAI V1 compatible API.
type Client interface {
	ListModels() (*ListModelResponse, error)
	Completion(request *CompletionReq) ([]CompletionResponse, error)
}

type client struct {
	baseURL string
	timeout uint // seconds
	client  *http.Client
	bearer  string
}

// NewClient creates a new OpenAI compatible V1 client.
//
// baseURL is the base URL of the API.
// timeout is the timeout for the HTTP requests in seconds.
// bearer is the bearer token for authentication.
func NewClient(baseURL string, timeout uint, bearer string) Client {
	return client{
		baseURL: baseURL,
		timeout: timeout,
		client: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
		bearer: bearer,
	}
}

// NewLMSClient creates a new OpenAI V1 compatible client
// for use with LM Studio with no bearer token.
//
// baseURL is the base URL of the API.
// timeout is the timeout for the HTTP requests in seconds.
func NewLMSClient(baseURL string, tineout uint) Client {
	return NewClient(baseURL, 10, "")
}
