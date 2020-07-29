package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		h.ServeHTTP(w, r)
		log.Printf("Request to '%s' took %dms", r.URL.String(), time.Since(startTime).Milliseconds())
		//log to a file if >20ms
	}
}
