package main

import (
	"fmt"
	"sync"
	"time"
)

var wait sync.WaitGroup
var counter int

func main() {
	wait.Add(2)
	defer wait.Wait()
	go add("first")
	go add("second")
}

func add(hint string) {
	for i := 0; i < 3; i++ {
		a := counter
		a++
		time.Sleep(time.Second * 2)
		counter = a
		fmt.Println("Hint = ", hint, " i = ", i, " count = ", counter)
	}
	wait.Done()
}
