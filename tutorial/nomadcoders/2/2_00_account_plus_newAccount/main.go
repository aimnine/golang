package main

import "fmt"
import "github.com/aimnine/golang/tutorial/nomadcoders/2/2_00_account_plus_newAccount/accounts"

func main(){
	account := accounts.NewAccount("aimine")
	fmt.Println(account)
}