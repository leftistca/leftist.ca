package litlenlib

import (
	"fmt"
	"path/filepath"
	"os"
	"log"
	"strings"
)

type Library struct {
	books []Book
	authors []Author
}

func NewLibrary(libraryPath string) (*Library, error) {
	lib := Library{
		books: []Book{},
		authors: []Author{},
	}

	authors, err := NewAuthorSlice(libraryPath + "authors.json")
	if(err != nil){
		log.Fatal(err)
	}
	lib.authors = authors

	//find all book.json files and parse them as books.
	err = filepath.Walk(libraryPath, func(path string, info os.FileInfo, err error) error {
		if( err != nil){
			return err
		}

		if(strings.ToLower(filepath.Base(path)) == "book.json"){ //if the file ends in 'book.json'
			book, err := NewBook(path, lib)
			if(err != nil){
				return err
			}
			
			lib.books = append(lib.books, book)
		}
		return nil
	})
	if(err != nil){
		log.Fatal(err)
	}

	return &lib, nil
}


func (lib Library) GetBook(id int) (Book, error) {
	for _, book := range lib.books {
		if(book.Id == id){
			return book, nil
		}
	}
	return Book{}, fmt.Errorf("Could not find book with id (%d)", id)
}

func (lib Library) GetAuthor(id int) (Author, error) {
	for _, author := range lib.authors {
		if(author.Id == id){
			return author, nil
		}
	}
	return Author{}, fmt.Errorf("Could not find author with id (%d)", id)
}