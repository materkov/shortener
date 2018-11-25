package redirecter

import (
	"context"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestApp_Create(t *testing.T) {
	s := NewStoreMemory()
	p := NewPubsub("")
	a := NewApp(s, p)

	id1, err := a.Create(context.Background(), "http://site1.com")
	assert.NoError(t, err)
	assert.NotZero(t, id1)

	id2, err := a.Create(context.Background(), "http://site2.com")
	assert.NoError(t, err)
	assert.NotZero(t, id2)

	id3, err := a.Create(context.Background(), "http://site3.com")
	assert.NoError(t, err)
	assert.NotZero(t, id3)

	assert.NotEqual(t, id1, id2)
	assert.NotEqual(t, id2, id3)
}

func TestApp_GetByKey(t *testing.T) {
	s := NewStoreMemory()
	p := NewPubsub("")
	a := NewApp(s, p)

	id1, err := s.Create(context.Background(), "http://site1.com")
	assert.NoError(t, err)

	id2, err := s.Create(context.Background(), "http://site2.com")
	assert.NoError(t, err)

	key1 := strconv.FormatInt(int64(id1), 36)
	key2 := strconv.FormatInt(int64(id2), 36)

	url1, err := a.GetByKey(context.Background(), key1)
	assert.NoError(t, err)
	assert.Equal(t, "http://site1.com", url1)

	url2, err := a.GetByKey(context.Background(), key2)
	assert.NoError(t, err)
	assert.Equal(t, "http://site2.com", url2)
}
