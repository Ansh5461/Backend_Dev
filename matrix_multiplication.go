package main

import "fmt"

/*arr1 := make([][]int, R)
	for i := 0; i < R; i++ {
		arr1[i] = make([]int, C)
	}
User defined array
*/

func main() {
	a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	b := [3][2]int{{1, 2}, {3, 4}, {5, 6}}

	c := [2][2]int{}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 3; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(c[i][j], "\t")
		}
		fmt.Println()
	}
}
