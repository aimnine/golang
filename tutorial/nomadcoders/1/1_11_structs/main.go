package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "ramen"}
	//	john := person{"john", 18, favFood}
	john := person{name: "john", age: 18, favFood: favFood}
	fmt.Println(john)
}
