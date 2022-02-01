package main

import "fmt"

func main() {
	fmt.Println(fac(5))
}

func fac(x int) int {
	if x == 1 || x == 0 {
		return 1
	}
	if x < 0 {
		return 0
	}
	r := x * fac(x-1)
	return r
}
