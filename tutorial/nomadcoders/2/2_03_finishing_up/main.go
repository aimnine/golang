package main

import (
	"fmt"
	"log"

	"github.com/aimnine/golang/tutorial/nomadcoders/2/2_02_methods_part_two/accounts"
)

func main() {
	account := accounts.NewAccount("aimnine")
	account.Deposit(100)
	fmt.Println(account.Balance())
	err := account.Withdraw(200)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())
}
