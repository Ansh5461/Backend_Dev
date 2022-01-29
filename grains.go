package main

import "fmt"

func main() {
	var c [64]uint64
	var s uint64
	c[0] = 1
	for i := 1; i < 64; i++ {
		c[i] = c[i-1] * 2
		s = s + c[i]
	}
	fmt.Println("Total grains are ", s)
}
