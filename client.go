package gobark

import "net/http"

var _ Dog = (*Client)(nil)

type Client struct {
	address string //https://xxx
	key     string //your auth key
	httpCli *http.Client
}

type Option func(c *Client)

func NewClient(address, key string, opts ...Option) *Client {
	c := &Client{
		address: "",
		key:     "",
		httpCli: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(c)
	}

	// 03/02/2023
	// TODO:
	return c
}

type CommonResp struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Data      any    `json:"data,omitempty"`
	Timestamp int64  `json:"timestamp"`
}

func (c *Client) Bark(req *BarkRequest) error {
	// 03/02/2023
	// TODO:
	return nil
}

func (c *Client) do(req, resp any) error {
	// 03/02/2023
	// TODO:
	return nil
}
