package redirecter

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	ErrBadBody  = NewError("bad_request", "Invalid request json")
	ErrInternal = NewError("internal_error", "Internal server error")
)

type Api struct {
	app  *App
	addr string
}

func NewAPI(app *App, addr string) *Api {
	return &Api{
		app:  app,
		addr: addr,
	}
}

func (a *Api) writeErr(w http.ResponseWriter, err error) {
	if _, ok := err.(*Error); ok {
		a.write(w, 400, err)
	} else {
		log.Printf("[ERROR] Internal error: %s", err)
		a.write(w, 500, ErrInternal)
	}
}

func (a *Api) write(w http.ResponseWriter, code int, resp interface{}) {
	a.writeCORS(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(resp)
}

func (a *Api) redirect(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["id"]
	url, err := a.app.GetByKey(r.Context(), key)
	if err != nil {
		// no json here
		w.WriteHeader(400)
		fmt.Fprintf(w, "Not existing URL %s", key)
		return
	}

	http.Redirect(w, r, url, 302)
}

type createRequest struct {
	URL string `json:"url"`
}

type createResponse struct {
	Key string `json:"key"`
}

func (a *Api) create(w http.ResponseWriter, r *http.Request) {
	req := createRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		a.writeErr(w, ErrBadBody)
		return
	}

	key, err := a.app.Create(r.Context(), req.URL)
	if err != nil {
		a.writeErr(w, err)
		return
	}

	a.write(w, 200, createResponse{Key: key})
}

func (a *Api) Serve() {
	r := mux.NewRouter()
	r.Methods("GET").Path("/redirect/{id}").HandlerFunc(a.redirect)
	r.Methods("POST").Path("/create").HandlerFunc(a.create)

	log.Printf("Starting HTTP server at %s", a.addr)
	err := http.ListenAndServe(a.addr, a.corsMiddleware(r))
	log.Fatal(err)
}
