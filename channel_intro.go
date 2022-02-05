package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan string)

	go func() {
		channel <- "Hello mf from channel"
		time.Sleep(time.Second * 5)
		fmt.Println(1)

	}()

	message := <-channel
	//time.Sleep(time.Second * 5)
	fmt.Println(message)

}

/*
func main() {

	channel := make(chan string, 1)

	go func() {
		channel <- "Hello mf from channel"
		channel <- "Hello again"
		fmt.Println(1)

	}()

	message := <-channel
	fmt.Println(message)

}
*/
