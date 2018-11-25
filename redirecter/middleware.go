package redirecter

import "net/http"

func (a *Api) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			a.writeCORS(w)
			w.WriteHeader(204)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (a *Api) writeCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
	w.Header().Set("Access-Control-Max-Age", "1728000")
}
