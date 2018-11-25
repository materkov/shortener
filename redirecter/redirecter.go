package redirecter

import (
	"context"
	"log"
	"net/url"
	"strconv"
)

const keyBase = 36

var (
	ErrBadURL = NewError("bad_url", "Invalid URL")
)

type App struct {
	store  Store
	pubsub *Pubsub
}

func NewApp(store Store, pubsub *Pubsub) *App {
	return &App{
		store:  store,
		pubsub: pubsub,
	}
}

func (a *App) GetByKey(ctx context.Context, key string) (url string, err error) {
	id, err := strconv.ParseInt(key, keyBase, 32)
	if err != nil {
		return "", ErrBadKey
	}

	url, err = a.store.GetByKey(ctx, int(id))
	if err != nil {
		return "", err
	}

	err = a.pubsub.Pub(topicClick, map[string]interface{}{"id": id})
	if err != nil {
		log.Printf("Error publishing msg: %s", err)
	}

	return url, nil
}

func (a *App) Create(ctx context.Context, urlAddr string) (key string, err error) {
	pUrl, err := url.ParseRequestURI(urlAddr)
	if err != nil {
		return "", ErrBadURL
	}
	log.Printf("%+v", pUrl)

	id, err := a.store.Create(ctx, urlAddr)
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(int64(id), keyBase), nil
}
