package libraryview

import (
	"html/template"
	"net/http"

	"../../models/litlenlib"
)

type HTMLLibraryView struct {
	template map[string]*template.Template
}

func NewHTMLLibraryView() *HTMLLibraryView {
	htmlview := HTMLLibraryView{
		template: map[string]*template.Template{},
	}
	htmlview.template["index"] = template.Must(template.ParseFiles("views/html_templates/index.html"))
	htmlview.template["bookinfo"] = template.Must(template.ParseFiles("views/html_templates/book.html"))
	htmlview.template["notfound"] = template.Must(template.ParseFiles("views/html_templates/notfound.html"))

	return &htmlview
}

func (view *HTMLLibraryView) RenderIndex(w http.ResponseWriter, books []*litlenlib.Book) error {
	w.WriteHeader(http.StatusOK)
	return view.template["index"].Execute(w, books)
}

func (view *HTMLLibraryView) RenderBookInfo(w http.ResponseWriter, book *litlenlib.Book) error {
	w.WriteHeader(http.StatusOK)
	return view.template["bookinfo"].Execute(w, book)
}

func (view *HTMLLibraryView) RenderNotFound(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusNotFound)
	return view.template["notfound"].Execute(w, nil)
}

/*
TODO:
	Proper Error Returning in NewHTMLLibraryView
*/
