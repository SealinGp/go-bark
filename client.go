package gobark

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

var _ Dog = (*Client)(nil)

const (
	defaultAddr = "https://api.day.app/"
)

type Client struct {
	addr    string
	key     string //your auth key
	httpCli *http.Client
	log     io.Writer
	debug   bool
}

func NewClient(key string, opts ...Option) (*Client, error) {
	c := &Client{
		addr:    defaultAddr,
		key:     key,
		httpCli: http.DefaultClient,
		log:     &emptyLog{},
	}

	for _, opt := range opts {
		opt(c)
	}

	if err := c.check(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) check() error {
	if c.addr == "" {
		return errors.New("addr required")
	}

	if c.key == "" {
		return errors.New("key required")
	}

	if c.httpCli == nil {
		return errors.New("http.Client required")
	}

	return nil
}

func (c *Client) Bark(ctx context.Context, req *BarkRequest) error {
	httpReq, err := c.newReq(ctx)
	if err != nil {
		c.fLog("[E] new req failed. err:%v", err)
		return err
	}

	req.set(httpReq)
	return c.do(ctx, httpReq)
}

func (c *Client) do(ctx context.Context, httpReq *http.Request, resp ...Responser) error {
	var r Responser = &CommonResp{}
	if len(resp) == 1 {
		r = resp[0]
	}

	httpResp, err := c.httpCli.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, r); err != nil {
		return err
	}

	if err = r.Error(); err != nil {
		return err
	}

	return nil
}

func (c *Client) newReq(ctx context.Context) (*http.Request, error) {
	basicURL, err := c.basicURL()
	if err != nil {
		c.fLog("[E] get basic URL failed. err:%v", err)
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, basicURL, nil)
	if err != nil {
		c.fLog("[E] new http request failed. url:%v, err:%v", basicURL, err)
		return nil, err
	}

	if c.debug {
		c.fLog("[D] %v", httpReq.URL)
	}

	return httpReq, nil
}

func (c *Client) basicURL() (string, error) {
	u, err := url.Parse(c.addr)
	if err != nil {
		c.fLog("[E] parse addr failed. addr:%v, err:%v", c.addr, err)
		return "", err
	}

	u.Path = fmt.Sprintf("/%v", c.key)
	return u.String(), nil
}

func (c *Client) fLog(format string, a ...interface{}) {
	l := fmt.Sprintf(format, a...)
	c.log.Write([]byte(l))
}
