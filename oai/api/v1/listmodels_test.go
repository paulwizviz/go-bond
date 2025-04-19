package oaiapi_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	oaiapi "github.com/paulwizviz/go-bond/oai/api/v1"
	"github.com/stretchr/testify/assert"
)

func TestListModel(t *testing.T) {
	testcases := []struct {
		name    string
		handler func(rw http.ResponseWriter, req *http.Request)
		want    *oaiapi.ListModelResponse
	}{
		{
			name: "ListModels",
			handler: func(rw http.ResponseWriter, req *http.Request) {
				// Mock response for ListModels
				rw.WriteHeader(http.StatusOK)
				rw.Header().Set("Content-Type", "application/json")
				rw.Write([]byte(`{"data": [
					{
					"id": "gemma-2-2b-it",
					"object": "model",
					"owned_by": "organization_owner"
					},
					{
					"id": "text-embedding-nomic-embed-text-v1.5",
					"object": "model",
					"owned_by": "organization_owner"
					}
				],
  				"object": "list"}`))
			},
			want: &oaiapi.ListModelResponse{
				Data: []oaiapi.ListModelsDataResponse{
					{
						ID:      "gemma-2-2b-it",
						Object:  "model",
						OwnedBy: "organization_owner",
					},
					{
						ID:      "text-embedding-nomic-embed-text-v1.5",
						Object:  "model",
						OwnedBy: "organization_owner",
					},
				},
				Object: "list",
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new HTTP server
			server := httptest.NewServer(http.HandlerFunc(tc.handler))
			defer server.Close()

			// Create a new client
			client := oaiapi.NewLMSClient(server.URL, 10)

			// Call the ListModels method
			resp, err := client.ListModels()
			if assert.NoError(t, err) {
				assert.Equal(t, tc.want, resp, fmt.Sprintf("expected %v, got %v", tc.want, resp))
			}
		})
	}
}
