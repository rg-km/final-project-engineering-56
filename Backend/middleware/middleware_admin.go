package middleware

import (
	"encoding/json"
	"net/http"
)

type Respon struct {
	Error string `json:"error"`
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do something before request
		next.ServeHTTP(w, r)
		// do something after request

		encoder := json.NewEncoder(w)
		role := r.Context().Value("role")
		if role != "admin" {
			w.WriteHeader(http.StatusUnauthorized)
			encoder.Encode(Respon{Error: "You are not admin"})
			return
		}
		next.ServeHTTP(w, r)
	})
}
