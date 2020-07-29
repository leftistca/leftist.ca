package libraryview

import (
	"encoding/json"
	"net/http"

	"../../models/litlenlib"
)

type JsonLibraryView struct{}

//TODO: Proper err checking
func NewJsonLibraryView() *JsonLibraryView {
	view := JsonLibraryView{}
	return &view
}

func (view *JsonLibraryView) RenderIndex(w http.ResponseWriter, books []*litlenlib.Book) error {
	response, err := json.Marshal(books)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)
	return err
}

func (view *JsonLibraryView) RenderBookInfo(w http.ResponseWriter, book *litlenlib.Book) error {
	response, err := json.Marshal(book)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)
	return err
}

func (view *JsonLibraryView) RenderNotFound(w http.ResponseWriter) error {
	response, err := json.Marshal(struct {
		Id int `json:"error"`
	}{404})
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	_, err = w.Write(response)
	return err
}
