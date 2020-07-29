package middleware

import (
	"net/http"
	"strings"
)

func AddTrailingSlash(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") == false {
			r.URL.Path = r.URL.Path + "/"
		}
		h.ServeHTTP(w, r)
	})
}
