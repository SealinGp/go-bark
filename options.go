package gobark

import (
	"io"
	"net/http"
)

type emptyLog struct{}

func (l *emptyLog) Write(p []byte) (n int, err error) {
	return
}

type Option func(c *Client)

func WithDebug() Option {
	return func(c *Client) {
		c.debug = true
	}
}

func WithAddr(addr string) Option {
	return func(c *Client) {
		c.addr = addr
	}
}

func WithHttpClient(httpCli *http.Client) Option {
	return func(c *Client) {
		c.httpCli = httpCli
	}
}

func WithLog(l io.Writer) Option {
	return func(c *Client) {
		c.log = l
	}
}
