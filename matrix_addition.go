package main

import "fmt"

func main() {
	a := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	/*fmt.Println("Enter your numbers")
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Scanf("%d", &a[i][j])
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d \t", a[i][j])
		}
		fmt.Println()
	}*/
	c := [3][2]int{}
	b := [3][2]int{{1, 2}, {3, 4}, {5, 6}}

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = a[i][j] + b[i][j]
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d\t", c[i][j])
		}
		fmt.Println()
	}
}
