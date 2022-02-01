package main

import "fmt"

var a = 5

func demo() {
	var a = 7
	fmt.Println(a)
}

func main() {
	demo()
	fmt.Println(a)
}
