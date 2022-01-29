package main

import "fmt"

type Macbook struct {
	processor int
	ram       int
	name      string
	price     int
	camera    int
	other     HP
}
type HP struct {
	price  int
	better string
}

func (i Macbook) call(name string) {
	fmt.Println(i.camera)
	fmt.Println(name)
}

func main() {
	a := Macbook{}
	a.processor = 123
	a.ram = 78
	a.name = "Ansh"
	a.other.price = 48
	fmt.Println(a.other.price)
	fmt.Println(a.ram)
}
