package main

import (
	"fmt"

	"github.com/aimnine/golang/tutorial/nomadcoders/2/2_04_dictionary_part_one/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	//	fmt.Println(dictionary["first"])
	//	definition, err := dictionary.Search("second")
	definition, err := dictionary.Search("first")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
