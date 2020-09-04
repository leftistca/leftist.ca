package middleware

import (
	"net/http"
	"sync"

	//"io/ioutil"
	"bytes"
	"fmt"
)

/* CUSTOM RESPONSEREADWRITER */

type ResponseReadWriter struct {
	body       bytes.Buffer
	statusCode int
	header     http.Header
}

func NewResponseReadWriter() *ResponseReadWriter {
	return &ResponseReadWriter{
		header: http.Header{},
	}
}

func (r *ResponseReadWriter) Header() http.Header {
	return r.header
}

func (r *ResponseReadWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return 0, nil
}

func (r *ResponseReadWriter) WriteHeader(statusCode int) {
	r.statusCode = statusCode
}

/* END CUSTOM RESPONSEREADWRITER */

type ResponseCache struct {
	responses map[string]ResponseReadWriter
	mux       sync.Mutex
}

func NewResponseCache() *ResponseCache {
	rc := ResponseCache{
		responses: map[string]ResponseReadWriter{},
	}
	return &rc
}

func (rc *ResponseCache) RecordResponse(h http.HandlerFunc, w http.ResponseWriter, r *http.Request) ResponseReadWriter {
	responseReadWriter := NewResponseReadWriter()
	h.ServeHTTP(responseReadWriter, r)

	rc.mux.Lock()
	rc.responses[r.URL.String()] = *responseReadWriter
	rc.mux.Unlock()

	return *responseReadWriter
}

func (rc *ResponseCache) HandleRequest(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		rc.mux.Lock()
		response, responseFound := rc.responses[r.URL.String()]
		rc.mux.Unlock()

		if !responseFound {
			fmt.Println("Request not cached. Caching it...")
			response = rc.RecordResponse(h, w, r)
		}

		fmt.Println("Serving cached request...")

		//copy all values in header to new response:
		for key, values := range response.Header() {
			for _, value := range values {
				w.Header().Set(key, value)
			}
		}

		w.WriteHeader(response.statusCode) //Set the status code
		//This call must come AFTER w.Header.Set(...) or else it seems like those values get overwritten.

		response.body.WriteTo(w) // Copy body to new ResponseWriter
	}
}
