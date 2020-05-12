package libraryController

import (
	"log"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	"leftist/views/libraryview"
	"leftist/models/litlenlib"
	"leftist/controllers/middleware"
)

func Initialize(prefix string, view libraryview.LibraryView, router *mux.Router, lib *litlenlib.Library,  responseCache *middleware.ResponseCache) (error){
	router.HandleFunc(prefix + "/book/{id}", middleware.LoggingMiddleware(responseCache.HandleRequest(BookInfo(view, lib))))
	router.HandleFunc(prefix + "/", middleware.LoggingMiddleware(NotFoundHandler(view)))
	return nil
}

func NotFoundHandler(v libraryview.LibraryView) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := v.RenderNotFound(w)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BookInfo(view libraryview.LibraryView, library *litlenlib.Library) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if(err != nil){
			NotFoundHandler(view)(w, r)
			return
		}
	
		book, err := library.GetBook(id)
		if(err != nil){
			NotFoundHandler(view)(w, r)
			return
		}
		err = view.RenderBookInfo(w, &book)
		if err != nil {
			log.Fatal(err)
		}
	}
}