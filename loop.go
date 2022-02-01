package main

import "fmt"

func main() {
	i := 1
	for i <= 5 {
		fmt.Println("Yo buddy")
		i++
	}

	for j := 1; j <= 5; j++ {
		fmt.Println("Its fucking amazing", j)
	}

	for j := 0; j < 3; j++ {
		for j := 0; j < 3; j++ {
			fmt.Println(j)
		}
	}

}
