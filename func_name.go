package main

import (
	"errors"
	"fmt"
)

func name(n string) (string, error) {
	if n == "" {
		return n, errors.New("name cant be empty")
	}
	return "Hello " + n, nil
}

func add(num ...int) int {
	sum := 0
	for _, value := range num {
		sum = sum + value
	}
	return sum

}

func main() {
	nam, err := name("")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(nam)
	}
	fmt.Println(add(1, 2, 3, 4, 5, 6))
}
