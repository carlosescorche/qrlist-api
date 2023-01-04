package middlewares

import (
	"net/http"
)

func MiddlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origins := r.Header["Origin"]

		if len(origins) == 0 {
			next.ServeHTTP(w, r)
			return
		}

		h := w.Header()
		h.Add("Access-Control-Allow-Credentials", "true")
		h.Add("Access-Control-Allow-Headers", "Content-Type")
		h.Add("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		h.Add("Access-Control-Allow-Origin", "*")

		next.ServeHTTP(w, r)
	})
}
