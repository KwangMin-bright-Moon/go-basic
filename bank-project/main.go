package main

import (
	"fmt"

	"github.com/KwangMin-bright-Moon/learngo/bank-project/banking"
)

func main(){
	account :=  banking.Account{Owner: "jason", Balance: 1000}
	fmt.Println(account)
}