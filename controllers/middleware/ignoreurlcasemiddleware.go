package middleware

import (
	"net/http"
	"strings"
)

func IgnoreURLCaseMiddleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = strings.ToLower(r.URL.Path)
        h.ServeHTTP(w, r)
    })
}