package main

import "fmt"

/*
switch(data){
case var1:
	statement
	no break statement

case var2:
	statements
	and continue......

case var3:
	statements
	fallthrough   for executing all statements below it

default:
	statements
}
*/
func main() {
	fmt.Println("Enter  a number")
	var inp int
	fmt.Scanln(&inp)
	if inp%2 == 0 {
		fmt.Println("Even buddy")
	} else {
		fmt.Println("Odd buddy")
	}

	x := 1000
	if x = 10; 100%2 == 0 {
		fmt.Println("Scope works buddy")
	} else {
		fmt.Println("Doesnt work buddy")
	}
	fmt.Println(x)

	i := 100
	switch i {
	case 10:
		i = 44
		fmt.Println(i)

	case 100:
		i = 60
		fmt.Println(i)
		fallthrough

	case 48:
		fmt.Println("Yo mf")
	}
}
