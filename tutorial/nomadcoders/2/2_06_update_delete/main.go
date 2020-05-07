package main

import (
	"fmt"

	"github.com/aimnine/golang/tutorial/nomadcoders/2/2_06_update_delete/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err := dictionary.Update(baseWord, "Second")
	if err != nil {
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)
	//	err2 := dictionary.Delete(baseWord)
	err2 := dictionary.Delete(baseWord + "asdf")
	fmt.Println(err2)
}
