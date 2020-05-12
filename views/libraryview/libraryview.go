package libraryview

import (
	"net/http"
	"leftist/models/litlenlib"
)
type LibraryView interface {
	RenderBookInfo(w http.ResponseWriter, book *litlenlib.Book) error
	RenderNotFound(w http.ResponseWriter) error
}