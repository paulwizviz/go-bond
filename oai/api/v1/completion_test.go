package oaiapi_test

import (
	"encoding/json"
	"fmt"
	"testing"

	oaiapi "github.com/paulwizviz/go-bond/oai/api/v1"
	"github.com/stretchr/testify/assert"
)

func TestCompletionRequest(t *testing.T) {
	tcs := []struct {
		request *oaiapi.CompletionReq
		want    []byte
	}{
		{
			request: func() *oaiapi.CompletionReq {
				prompt := oaiapi.PromptType{
					Kind:        oaiapi.PromptKindString,
					StringValue: "Hello, world!",
				}
				return &oaiapi.CompletionReq{
					Model:  "gpt-3.5-turbo",
					Prompt: prompt,
				}
			}(),
			want: []byte(`{"model":"gpt-3.5-turbo","prompt":"Hello, world!"}`),
		},
		{
			request: func() *oaiapi.CompletionReq {
				prompt := oaiapi.PromptType{
					Kind:        oaiapi.PromptKindStringArray,
					StringArray: []string{"Hello", "world"},
				}
				return &oaiapi.CompletionReq{
					Model:  "gpt-3.5-turbo",
					Prompt: prompt,
				}
			}(),
			want: []byte(`{"model":"gpt-3.5-turbo","prompt":["Hello","world"]}`),
		},
		{
			request: func() *oaiapi.CompletionReq {
				prompt := oaiapi.PromptType{
					Kind:       oaiapi.PromptKindTokenArray,
					TokenArray: []oaiapi.Token{"Hello", "world"},
				}
				return &oaiapi.CompletionReq{
					Model:  "gpt-3.5-turbo",
					Prompt: prompt,
				}
			}(),
			want: []byte(`{"model":"gpt-3.5-turbo","prompt":["Hello","world"]}`),
		},
		{
			request: func() *oaiapi.CompletionReq {
				prompt := oaiapi.PromptType{
					Kind: oaiapi.PromptKindArrayOfTokens,
					ArrayOfTokens: [][]oaiapi.Token{
						{"hello", "world"},
						{"ola", "world"},
					},
				}
				return &oaiapi.CompletionReq{
					Model:  "gpt-3.5-turbo",
					Prompt: prompt,
				}
			}(),
			want: []byte(`{"model":"gpt-3.5-turbo","prompt":[["hello","world"],["ola","world"]]}`),
		},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			b, err := json.Marshal(tc.request)
			if assert.NoError(t, err) {
				assert.Equal(t, string(tc.want), string(b), "unexpected json")
			}
		})
	}
}

func TestCompletionResponseUnmarshal(t *testing.T) {
	tcs := []struct {
		data []byte
		want oaiapi.CompletionResponse
	}{
		{
			data: []byte(`{"id":"cmpl-m7c8gu3wxv789zszjqus",
			"object":"text_completion",
			"created":1745004685,
			"model":"gemma-2-2b-it",
			"choices":[
			   {"index":0,
				"text":"\n\n<ul>\n  <li>Bitcoin</li>\n  <li>Ethereum</li>\n  <li>Polygon</li>\n</ul>\n\n\nPlease note: This is for educational purposes only and should not be taken as financial advice. \n",
				"logprobs":null,
				"finish_reason":"stop"}
			],
			"usage":{
				"prompt_tokens":19,
				"completion_tokens":56,
				"total_tokens":75
			},
			"stats":{}}`),
			want: func() oaiapi.CompletionResponse {
				return oaiapi.CompletionResponse{
					ID:      "cmpl-m7c8gu3wxv789zszjqus",
					Object:  "text_completion",
					Created: 1745004685,
					Model:   "gemma-2-2b-it",
					Choices: []oaiapi.CompletionChoicesResp{
						{
							Index:        0,
							Text:         "\n\n<ul>\n  <li>Bitcoin</li>\n  <li>Ethereum</li>\n  <li>Polygon</li>\n</ul>\n\n\nPlease note: This is for educational purposes only and should not be taken as financial advice. \n",
							LogProbs:     nil,
							FinishReason: "stop",
						},
					},
					Usage: oaiapi.CompletionUsageResp{
						PromptTokens:     19,
						CompletionTokens: 56,
						TotalTokens:      75,
					},
				}
			}(),
		},
		{
			data: []byte(`{"id":"cmpl-m7c8gu3wxv789zszjqus",
			"object":"text_completion",
			"created":1745004685,
			"model":"gemma-2-2b-it",
			"choices":[
			   {"index":0,
			    "text":"\n\n<ul>\n  <li>Bitcoin</li>\n  <li>Ethereum</li>\n  <li>Polygon</li>\n</ul>\n\n\nPlease note: This is for educational purposes only and should not be taken as financial advice. \n",
				"logprobs": {
                    "text_offset": [0, 6, 12],
                    "token_logprobs": [-0.1, -0.2, -0.05],
                    "tokens": ["The", " quick", " brown"],
                    "top_logprobs": [
                        {"The": -0.1, "A": -1.2, "An": -1.5},
                        {" quick": -0.2, " fast": -0.4},
                        {" brown": -0.05, " red": -0.3}
					]
                },
				"finish_reason":"stop"}],
			"usage":{
				"prompt_tokens":19,
				"completion_tokens":56,
				"total_tokens":75
			},
			"stats":{}}`),
			want: func() oaiapi.CompletionResponse {
				return oaiapi.CompletionResponse{
					ID:      "cmpl-m7c8gu3wxv789zszjqus",
					Object:  "text_completion",
					Created: 1745004685,
					Model:   "gemma-2-2b-it",
					Choices: []oaiapi.CompletionChoicesResp{
						{
							Index: 0,
							Text:  "\n\n<ul>\n  <li>Bitcoin</li>\n  <li>Ethereum</li>\n  <li>Polygon</li>\n</ul>\n\n\nPlease note: This is for educational purposes only and should not be taken as financial advice. \n",
							LogProbs: &oaiapi.CompletionLogProbsResp{
								TextOffset:    []int{0, 6, 12},
								TokenLogProbs: []float64{-0.1, -0.2, -0.05},
								Tokens:        []string{"The", " quick", " brown"},
								TopLogProbs:   []map[string]float64{{"The": -0.1, "A": -1.2, "An": -1.5}, {" quick": -0.2, " fast": -0.4}, {" brown": -0.05, " red": -0.3}},
							},
							FinishReason: "stop",
						},
					},
					Usage: oaiapi.CompletionUsageResp{
						PromptTokens:     19,
						CompletionTokens: 56,
						TotalTokens:      75,
					},
				}
			}(),
		},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			var got oaiapi.CompletionResponse
			err := json.Unmarshal(tc.data, &got)
			if assert.NoError(t, err) {
				assert.Equal(t, tc.want, got, "unexpected response")
			}
		})
	}
}
