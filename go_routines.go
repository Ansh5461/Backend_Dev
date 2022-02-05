package main

import (
	"fmt"
	"sync"
	"time"
)

/*
this is for anonymous function without any mutex
func main() {
	//var mutex sync.Mutex
	counter := 0
	for i := 0; i < 10; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second * 5)
	fmt.Println(counter)
}
*/

/*
this is for anonymous function with mutex, still it is giving same output
func main() {
	var mutex sync.Mutex
	counter := 0
	for i := 0; i < 10; i++ {
		go func() {
			mutex.Lock()
			counter++
			mutex.Unlock()
		}()
	}
	//time.Sleep(time.Second * 5)
	fmt.Println(counter)
}
*/

func main() {
	var mutex sync.Mutex
	counter := 0
	for i := 0; i < 10; i++ {
		go func(i int) {
			mutex.Lock()
			counter++
			fmt.Println(i, " => ", counter)
			mutex.Unlock()
		}(i)
	}
	time.Sleep(time.Second * 5)
	fmt.Println(counter)
}
