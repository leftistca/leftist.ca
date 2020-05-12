package main

import (
	"log"
	"net/http"
	"leftist/views/libraryview"
	"leftist/models/litlenlib"
	"leftist/controllers"
	"github.com/gorilla/mux"
	"leftist/controllers/middleware"
)

func main() {

	library, err := litlenlib.NewLibrary("library_content/")
	if(err != nil){
		log.Fatal(err)
	}

	htmlLibraryView := libraryview.NewHTMLLibraryView()
	jsonLibraryView := libraryview.NewJsonLibraryView()

	router := mux.NewRouter()
	responseCache := middleware.NewResponseCache()

	libraryController.Initialize("/library", htmlLibraryView, router, library, responseCache)
	libraryController.Initialize("/api", jsonLibraryView, router, library, responseCache)

	log.Fatal(http.ListenAndServe(":8080", middleware.IgnoreURLCaseMiddleware(router)))
}

/* TODO: Figure out why CORS doesn't work on succesful /api calls.
	https://stackoverflow.com/questions/40985920/making-golang-gorilla-cors-handler-work
*/