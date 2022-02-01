package main

import "fmt"

// defer stores the value in stack

func main() {
	a()

	// now lets practice multiple defers
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("hello motherfucker")

}
func a() {
	fmt.Println("a begins")
	defer b()
	// defer command executes b after execution of a, means not even in a, after a
	// it will evaluate b, but it will execute it only after full execution from where it is called
	fmt.Println("a ends")
}
func b() {
	fmt.Println("in b")
}
