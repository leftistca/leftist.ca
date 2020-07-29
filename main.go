package main

import (
	"log"
	"net/http"

	libraryController "./controllers"
	"./controllers/middleware"
	"./models/litlenlib"
	"./views/libraryview"

	"github.com/gorilla/mux"
)

func main() {

	library, err := litlenlib.NewLibrary("library_content/")
	if err != nil {
		log.Fatal(err)
	}

	htmlLibraryView := libraryview.NewHTMLLibraryView()
	jsonLibraryView := libraryview.NewJsonLibraryView()

	router := mux.NewRouter()
	responseCache := middleware.NewResponseCache()

	libraryController.Initialize("", htmlLibraryView, router, library, responseCache)
	libraryController.Initialize("/api", jsonLibraryView, router, library, responseCache)

	log.Fatal(http.ListenAndServe(":8080", middleware.AddTrailingSlash(middleware.IgnoreURLCase(router))))
}

/* TODO: Figure out why CORS doesn't work on succesful /api calls.
https://stackoverflow.com/questions/40985920/making-golang-gorilla-cors-handler-work
*/
