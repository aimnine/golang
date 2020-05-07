package main

import (
	"fmt"

	"github.com/aimnine/golang/tutorial/nomadcoders/2/2_01_methods_part_one/accounts"
)

func main() {
	account := accounts.NewAccount("aimnine")
	account.Deposit(100)
	fmt.Println(account.Balance())
}
