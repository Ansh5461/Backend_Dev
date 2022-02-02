package main

import "fmt"

func main() {
	fmt.Println(fibo(10))
}
func fibo(n int) int {
	if n == 0 {
		return n
	}
	if n == 1 {
		return n
	}
	res := fibo(n-1) + fibo(n-2)
	return res
}
