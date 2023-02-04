package gobark

import (
	"context"
	"log"
	"os"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	key := os.Getenv("GOBARK_KEY")

	mockReqs := []struct {
		name          string
		input         func() (*Client, error)
		doAndValidate func(*Client, error) error
	}{
		// {
		// 	name: "TextMsg",
		// 	input: func() (*Client, error) {
		// 		return NewClient(key)
		// 	},
		// 	doAndValidate: func(c *Client, err error) error {
		// 		if err != nil {
		// 			return err
		// 		}

		// 		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		// 		defer cancel()

		// 		return c.Bark(ctx, &BarkRequest{
		// 			Text: &Text{
		// 				Title:   "title",
		// 				Content: "content",
		// 			},
		// 		})
		// 	},
		// },
		{
			name: "SoundMsg",
			input: func() (*Client, error) {
				return NewClient(key, WithDebug(), WithLog(log.Writer()))
			},
			doAndValidate: func(c *Client, err error) error {
				if err != nil {
					return err
				}

				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()

				return c.Bark(ctx, &BarkRequest{
					BarkRequestOptions: BarkRequestOptions{
						Sound: "minuet",
					},
					Text: &Text{
						Title: "推送铃声",
					},
				})
			},
		},
	}

	for _, mockReq := range mockReqs {
		t.Run(mockReq.name, func(t *testing.T) {
			c, err := mockReq.input()

			if err = mockReq.doAndValidate(c, err); err != nil {
				t.Errorf("do and validate failed. err:%v", err)
				return
			}
		})
	}
}
