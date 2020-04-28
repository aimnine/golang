package main

import "fmt"

func main() {
	//	names := [5]string{"a", "b", "c"} // array
	//	names[3] = "d"
	//	names[4] = "e"

	names := []string{"a", "b", "c"} // slice
	names = append(names, "d")

	fmt.Println(names)
}
