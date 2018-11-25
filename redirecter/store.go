package redirecter

import (
	"context"
)

var ErrBadKey = NewError("bad_key", "Bad short key")

type Store interface {
	GetByKey(ctx context.Context, id int) (url string, err error)
	Create(ctx context.Context, url string) (id int, err error)
}
