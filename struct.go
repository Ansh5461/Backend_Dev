package main

import "fmt"

type Student struct {
	rollno int
	name   string
	marks  int
}

func main() {
	var s1 = Student{45, "Ansh", 100}
	fmt.Println(s1)
	fmt.Println(s1.name)
}
