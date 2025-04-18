package oaiapi

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type ListModelsDataResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	OwnedBy string `json:"owned_by"`
}

type ListModelResponse struct {
	Data   []ListModelsDataResponse `json:"data"`
	Object string
}

func (o client) ListModels() (*ListModelResponse, error) {
	endPoint, err := url.JoinPath(o.baseURL, "/v1/models")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, endPoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	response, err := o.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		io.Copy(io.Discard, response.Body) // Ensure the body is fully read
		response.Body.Close()
	}()

	decoder := json.NewDecoder(response.Body)

	var lrm ListModelResponse
	if err := decoder.Decode(&lrm); err != nil {
		return nil, err
	}

	return &lrm, nil
}
