package siftscience

import (
	"github.com/dghubble/sling"
)

var (
	Production = Environment{
		Url: "https://api.siftscience.com/",
		Key: "",
	}

	Sandbox = Environment{
		Url: "https://api.siftscience.com/",
		Key: "9e68c1b4132a432b",
	}
)

const (
	Name = "siftscience"
)

// Siftscience environment
type Environment struct {
	Url, Key string
}

// Siftscience client
type Client struct {
	Api *sling.Sling
	Env Environment
}

// Siftscience query params
type Params struct {
	ApiKey     string `url:"api_key,omitempty"`
	AbuseTypes string `url:"abuse_types,omitempty"`
}

// New Siftscience API client
func New(env Environment) *Client {
	api := sling.New().Client(nil)
	api.Add("User-Agent", "Ingresse-Payment/1.0")

	client := Client{
		Api: api.New().Base(env.Url),
		Env: env,
	}

	return &client
}

// Get user score
func (sift *Client) GetScore(id string) (*Score, error) {
	score := new(Score)

	params := &Params{
		ApiKey:     sift.Env.Key,
		AbuseTypes: "payment_abuse,account_abuse",
	}

	_, err := sift.Api.Get("v204/score/" + id).QueryStruct(params).ReceiveSuccess(score)

	if score.IsOk() {
		return score, nil
	}

	return nil, err
}
