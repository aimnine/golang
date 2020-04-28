package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	len, upper := lenAndUpper("john")
	fmt.Println(len, upper)
}
