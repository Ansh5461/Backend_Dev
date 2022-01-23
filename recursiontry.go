package main

import "fmt"

func main() {
	fmt.Println(f(5))
	fmt.Print("\n")
	x(4)
}
func f(n int) int {
	if n <= 0 {
		return 0
	}
	if n%2 == 0 {
		fmt.Println(n + 1)
		return f(n - 1)
	} else {
		fmt.Println(n + 2)
		return f(n - 1)
	}
}

func x(n int) {
	if n <= 0 {
		return
	}
	x(n - 1)
	x(n - 2)
	fmt.Println(n - 1)
}
