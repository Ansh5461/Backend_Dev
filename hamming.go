package main

import (
	"fmt"
)

func main() {
	var g1 string
	g1 = "GAGCCTACTAACGGGAT"
	var g2 string
	g2 = "CATCGTAATGACGGCCT"
	var c int
	for i := range g1 {
		if g1[i] != g2[i] {
			c = c + 1
		}
	}
	fmt.Println(c)
}
