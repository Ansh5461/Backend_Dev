package main

import (
	"fmt"
	"sync"
)

/*
var wait sync.WaitGroup

func main() {
	wait.Add(2)
	go hello()
	go hello()
	wait.Wait()
}

func hello() {
	fmt.Println("hello")
	wait.Done()
}*/

func main() {
	var wait sync.WaitGroup

	counter := 200
	wait.Add(counter)

	for i := 0; i < counter; i++ {
		go hello(&wait, i)
	}
	defer wait.Wait()
}

func hello(wait *sync.WaitGroup, counter int) {
	fmt.Print(counter, " ")
	wait.Done()

}
