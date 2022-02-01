package main

import "fmt"

func main() {
	i := 10
	//var j int  different ways of creating pointers
	//j := new(int)   // this is declaration + assignment (to some random free address) compared to nil by upper addresss
	j := &i
	fmt.Println(i)
	*j = 100
	fmt.Println(i)
	fmt.Println(j)

	name := new(string)
	*name = "ansh"
	fmt.Println(*name)
}
