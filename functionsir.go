package main

import "fmt"

func main() {
	w := new(int)
	name := new(string)
	result, x := c(w, name)
	fmt.Println(x)
	fmt.Println(result)
	fmt.Println(*w)
	fmt.Println(*name)
	b(1, 2, 3, 4, 5, 6)
}

func a() (int, string) {
	return 123, "Ansh"
}

func b(args ...int) {
	for _, value := range args {
		fmt.Print(value)
	}

}

func c(p *int, q *string) (i int, j string) {
	i = 10
	j = "abcd"
	*p = 69
	*q = "Itachi uchiha"
	return
}
