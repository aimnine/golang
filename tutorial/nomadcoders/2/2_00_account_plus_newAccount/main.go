package main

import (
	"fmt"
	"github.com/aimnine/golang/tutorial/nomadcoders/2/2_00_account_plus_newAccount/test"
	"github.com/aimnine/golang/tutorial/nomadcoders/2/2_00_account_plus_newAccount/accounts"
)

func main(){
	account := accounts.NewAccount("aimine")
	fmt.Println(account)
	t := test.Test("asdf")
	fmt.Println(t)
}