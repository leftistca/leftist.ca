package libraryview

import (
	"net/http"
	"encoding/json"
	"leftist/models/litlenlib"
)

type JsonLibraryView struct {}

//TODO: Proper err checking
func NewJsonLibraryView() *JsonLibraryView{
	view := JsonLibraryView{}
	return &view
}

func (view *JsonLibraryView) RenderBookInfo(w http.ResponseWriter, book *litlenlib.Book) error {
	response, err := json.Marshal(book)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return nil
}

func (view *JsonLibraryView) RenderNotFound(w http.ResponseWriter) error {
	response, err := json.Marshal(struct{Id int `json:"error"`}{404})
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(response)
	return nil
}