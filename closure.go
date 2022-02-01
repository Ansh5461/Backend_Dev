package main

import "fmt"

func main() {
	num := 10
	fmt.Println(num)

	var getInt func(int) int
	getInt = func(x int) int {
		fmt.Println("In closure")
		return 10 + x
	}
	// here we have defined a closure as getInt with func(int) int type, and we are passing that closure as a parameter
	// into the function, we are not sending getInt(100) or something because that would be of int type
	// and not func(int) int type
	res := g(getInt, 20)
	fmt.Println(res)

	j := func(i int) int {
		return i + 10
	}(10)

	fmt.Println(j)
}

func g(x func(int) int, y int) int {
	z := x(100)
	return z + y
}
