package main

import "fmt"

func submul(x int, y int) (int, int) {
	var z = x - y
	m := x * y
	return z, m
}

func add(x, y int) (out1, out2 int) {
	out1 = x * 2
	out2 = y * 2
	return
}

func main() {
	//default = 0
	var num1 int = 69
	var num2 int = 45

	var result = num1 + num2

	var num3, num4 int
	num3, num4 = 45, 82

	num5 := 100 //its shorthand for var num5 = 100

	const num6 = 72 //its value cant be changed, and if we try to, then it will not print anything

	res1, res2 := submul(num1, num2)
	name := "Ansh"
	a, b := add(num1, num2)
	fmt.Print(result, "\n")
	fmt.Print(num3 + num4)
	fmt.Print("\n", num5)
	fmt.Print("\n", res1, "\t", res2)
	fmt.Println("\n", a, b)
	fmt.Print(name)
}
