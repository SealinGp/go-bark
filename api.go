package gobark

import "context"

type Dog interface {
	Bark(ctx context.Context, req *BarkRequest) error
}
