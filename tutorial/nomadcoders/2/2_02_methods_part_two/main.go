package main

import (
	"fmt"
	"log"

	"github.com/aimnine/golang/tutorial/nomadcoders/2/2_03_finishing_up/accounts"
)

func main() {
	account := accounts.NewAccount("aimnine")
	account.Deposit(100)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}
	//	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)

}
