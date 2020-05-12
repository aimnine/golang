package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	// people := [2]string{"aimnine", "aimeight"}
	// for _, person := range people {
	// 	go isSexy(person, c)
	// }
	// time.Sleep(time.Second * 10)

	// result := <-c
	// fmt.Println(result)
	// fmt.Println("Waiting...")
	// resultOne := <-c
	// resultTwo := <-c
	// fmt.Println("Received this message: ", resultOne)
	// fmt.Println("Received this message: ", resultTwo)
	// fmt.Println(<-c) // deadlock!

	people := [5]string{"aimnine", "aimeight", "aimeseven", "aimsix", "aimfive"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println("waiting for", i)
		fmt.Println(<-c)
	}

}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is sexy"
}
