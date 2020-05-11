package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	people := [2]string{"aimnine", "aimeight"}
	for _, person := range people {
		go isSexy(person, c)
	}
	// time.Sleep(time.Second * 10)

	// result := <-c
	// fmt.Println(result)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c) // deadlock!
}

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- true
}
