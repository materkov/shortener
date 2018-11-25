package redirecter

import (
	"context"
	"math/rand"
)

type StoreMemory struct {
	urls map[int]string
}

func NewStoreMemory() Store {
	return &StoreMemory{
		urls: make(map[int]string),
	}
}

func (m *StoreMemory) GetByKey(ctx context.Context, id int) (url string, err error) {
	if url, ok := m.urls[id]; ok {
		return url, nil
	} else {
		return "", ErrBadKey
	}
}

func (m *StoreMemory) Create(ctx context.Context, url string) (id int, err error) {
	id = rand.Intn(2000000000)
	m.urls[id] = url
	return id, nil
}
