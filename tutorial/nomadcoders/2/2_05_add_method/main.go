package main

import (
	"fmt"

	"github.com/aimnine/golang/tutorial/nomadcoders/2/2_05_add_method/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	word := "hello"
	definition := "Greeting"
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
	hello, _ := dictionary.Search(word)
	fmt.Println(hello)
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
}
