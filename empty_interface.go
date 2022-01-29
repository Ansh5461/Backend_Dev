package main

import "fmt"

// empty interfaces are used for function overloading
//normal interface for polymorphism and abstraction

func describe(i interface{}) {
	//fmt.Printf("Type = %T Value = %v \n", i, i)
	switch i.(type) {
	case string:
		fmt.Println("String mf")
	case int:
		fmt.Println("Integer mf")
	case bool:
		fmt.Println("Boolean integer")
	default:
		fmt.Println("Unknown mf")
	}
}

func main() {
	describe("hello")
	describe(123)
	describe(12.34)
	describe(true)
}
