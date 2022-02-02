package main

import "fmt"

func main() {
	// now construct different instances of structs
	sa := SavingAccount{
		amount: 123,
	}
	na := NormalAccount{
		amount: 456,
	}

	// bc is an array of type balance_calculator which can handle all both type of data of na and sa
	// because it hs implemented the methods of SavingAccount and NormalAccount
	// so it can access the variables of both of them
	bc := []balance_calculator{sa, na}
	var total int
	for _, v := range bc {
		total = total + v.balance()
	}
	fmt.Println(total)
	fmt.Println(sa.amount)
	fmt.Println(na.amount)
}

type SavingAccount struct {
	amount int
}

func (sa SavingAccount) balance() int {
	return sa.amount
}

type NormalAccount struct {
	amount int
}

func (na NormalAccount) balance() int {
	return na.amount
}

type balance_calculator interface {
	balance() int
}
