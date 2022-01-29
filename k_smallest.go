package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Println("Enter the size of array")
	fmt.Scanf("%d", &n)

	arr := make([]int, n)
	fmt.Print("Enter numbers \t")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Print("Enter k, k should be less than ", n, "\t")
	var k int
	fmt.Scan(&k)

	// Input taken, now for implementation
	var ans int
	sort.Ints(arr)
	ans = arr[k-1]
	fmt.Println(k, " smallest number is ", ans)

}
