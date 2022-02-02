package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the size of array")
	fmt.Scanf("%d", &n)

	arr := make([]int, n)
	fmt.Print("Enter numbers \t")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	var k int
	fmt.Print("Enter the number whose frequency you want to find \t")
	fmt.Scan(&k)
	var c int
	for i := 0; i < n; i++ {
		if arr[i] == k {
			c = c + 1
		}
	}
	fmt.Println("Frequency = ", c)
}
