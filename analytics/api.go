package analytics

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Api struct {
	app    *App
	pubsub *Pubsub
	addr   string
}

func NewApi(app *App, pubsub *Pubsub, addr string) *Api {
	return &Api{app, pubsub, addr}
}

type onClickRequest struct {
	ID int `json:"id"`
}

func (a *Api) onClick(r []byte) {
	log.Printf("mq click: %s", r)

	req := onClickRequest{}
	if err := json.Unmarshal(r, &req); err != nil {
		return
	}

	_ = a.app.AddClick(context.Background(), req.ID)
}

func (a *Api) Serve() {
	go a.serveMQ()
	r := mux.NewRouter()

	err := http.ListenAndServe(a.addr, r)
	log.Fatal(err)
}

func (a *Api) serveMQ() {
	log.Printf("Starting listening mq")

	topics := []string{"click"}
	_ = a.pubsub.Listen(topics, "analytics", func(topic string, req []byte) {
		switch topic {
		case "click":
			a.onClick(req)
		}
	})
}

func (a *Api) serveHTTP() {
}
