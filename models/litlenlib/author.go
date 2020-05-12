package litlenlib

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type Author struct {
	Id int			`json:"id"`
	Name string		`json:"name"`
	Birthyear int	`json:"birthyear"`
}

func (a Author) Age() int {
	return 2020 - a.Birthyear
}

func NewAuthorSlice(path string) ([]Author, error) {
	authors := []Author{}

	author_file, err := ioutil.ReadFile(path)
	if(err != nil){
		return []Author{}, err
	}

	err = json.Unmarshal(author_file, &authors)
	if(err != nil){
		return []Author{}, err
	}

	//go through each author and ensure there are no id duplicates.
	for i, authorA := range authors {
		for _, authorB := range authors[i+1:] {
			if(authorA.Id == authorB.Id){
				return []Author{}, fmt.Errorf("Failed to parse authors.json, duplicate author ID (%d) present in (%s) and (%s).", authorA.Id, authorA.Name, authorB.Name)
			}
		}
	}

	return authors, nil
}