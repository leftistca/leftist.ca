package libraryview

import (
	"net/http"
	"html/template"
	"leftist/models/litlenlib"
)

type HTMLLibraryView struct {
	template map[string]*template.Template
}


func NewHTMLLibraryView() *HTMLLibraryView{
	htmlview := HTMLLibraryView{
		template: map[string]*template.Template{},
	}

	htmlview.template["bookinfo"] = template.Must(template.ParseFiles("views/html_templates/book.html"))
	htmlview.template["notfound"] = template.Must(template.ParseFiles("views/html_templates/notfound.html"))

	return &htmlview
}

func (view *HTMLLibraryView) RenderBookInfo(w http.ResponseWriter, book *litlenlib.Book) error {
	w.WriteHeader(http.StatusOK)
	view.template["bookinfo"].Execute(w, book)
	return nil
}

func (view *HTMLLibraryView) RenderNotFound(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusNotFound)
	view.template["notfound"].Execute(w, nil)
	return nil
}

/* 
TODO:
	Proper Error Returning in NewHTMLLibraryView
*/