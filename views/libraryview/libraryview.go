package libraryview

import (
	"net/http"

	"../../models/litlenlib"
)

type LibraryView interface {
	RenderIndex(w http.ResponseWriter, books []*litlenlib.Book) error
	RenderBookInfo(w http.ResponseWriter, book *litlenlib.Book) error
	RenderNotFound(w http.ResponseWriter) error
}
