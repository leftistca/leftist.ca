package libraryController

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"../controllers/middleware"

	"../models/litlenlib"
	"../views/libraryview"

	"github.com/gorilla/mux"
)

func Initialize(routePrefix string, view libraryview.LibraryView, router *mux.Router, lib *litlenlib.Library, responseCache *middleware.ResponseCache) error {
	router.HandleFunc(routePrefix+"/", middleware.Logging(Index(view, lib)))

	router.HandleFunc(routePrefix+"/book/{id}/", middleware.Logging(responseCache.HandleRequest(BookInfo(view, lib))))
	router.HandleFunc(routePrefix+"/unknown/", middleware.Logging(NotFoundHandler(view)))
	return nil
}

func Index(view libraryview.LibraryView, lib *litlenlib.Library) http.HandlerFunc {
	//get four random books, and display them on the home screen.
	return func(w http.ResponseWriter, r *http.Request) {
		randomBooks := []*litlenlib.Book{}

		for i := 0; i < 4; i++ {
			randomBook, err := lib.GetRandomBook()
			if err != nil {
				log.Fatal(err)
			}
			randomBooks = append(randomBooks, &randomBook)
		}
		fmt.Printf("%+v\n", randomBooks[0])
		err := view.RenderIndex(w, randomBooks)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func NotFoundHandler(view libraryview.LibraryView) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := view.RenderNotFound(w)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BookInfo(view libraryview.LibraryView, library *litlenlib.Library) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			NotFoundHandler(view)(w, r)
			return
		}

		book, err := library.GetBook(id)
		if err != nil {
			NotFoundHandler(view)(w, r)
			return
		}
		err = view.RenderBookInfo(w, &book)
		if err != nil {
			log.Fatal(err)
		}
	}
}
