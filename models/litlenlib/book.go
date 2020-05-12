package litlenlib

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Book struct {
	Id int				`json:"id"`
	Authors []*Author	//not automatically unmarshalled.
	Title string		`json:"title"`
	Year int			`json:"year"`
}

func NewBook(path string, library Library) (Book, error) {
	book_file, err := ioutil.ReadFile(path)
	if(err != nil){
		return Book{}, fmt.Errorf("Failed to create new Book: %v", err)
	}

	book := Book{}

	//Extract just the Authour IDs from the json
	authorIds := struct{Ids []int	`json:"author_ids"`}{}
	err = json.Unmarshal(book_file, &authorIds)
	if err != nil {
		return Book{}, err
	}

	//Extract the rest of the book's information from the Json
	err = json.Unmarshal(book_file, &book)
	if(err != nil){
		return Book{}, err
	}

	//Ensure book.id is unique
	_, err = library.GetBook(book.Id)
	if(err == nil){
		return Book{}, fmt.Errorf("Failed to parse '%s'. Book with id (%d) already exists.", book.Title, book.Id)
	}

	//Insert authors struct pointers from the previously extracted Authour IDs
	for _, authorId := range authorIds.Ids {
		author, err := library.GetAuthor(authorId)
		if(err != nil){
			return Book{}, fmt.Errorf("Failed to parse '%s'. Author with id (%d) does not exist", book.Title, authorId)
		}

		book.Authors = append(book.Authors, &author)
	}

	return book, nil
}
