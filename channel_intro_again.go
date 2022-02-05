package main

import (
	"fmt"
)

/*
func main() {
	helloch := make(chan string, 1)
	goodbyech := make(chan string, 1)
	quitch := make(chan bool)
	go receiveMessage(helloch, goodbyech, quitch)
	time.Sleep(time.Second)
	go sendMessage(helloch, "Hello mf")
	go sendMessage(goodbyech, "Goodbye mf")

	result := <-quitch
	fmt.Println("Message from quitch", result)

}

func sendMessage(ch chan<- string, message string) {
	ch <- message
}

func receiveMessage(helloch, goodbyech <-chan string, quitch chan<- bool) {
	for {
		select {
		case message := <-helloch:
			fmt.Println("Message from helloch = ", message)

		case message := <-goodbyech:
			fmt.Println("Message from goodbyech = ", message)

		case <-time.After(time.Second * 2):
			fmt.Println("Nothing receiving in 2 seconds, Exit")
			quitch <- true
			break
		}
	}
}*/

/*
func main() {
	channel := make(chan int)
	go func() {
		channel <- 1
		time.Sleep(time.Second)
		channel <- 2
		close(channel)
		//without this upper line, we will get a deadlock
	}()

	for i := range channel {
		fmt.Println(i)
	}
}
*/
func sendbool(ch chan bool) {
	ch <- true
}

func sendint(ch chan int) {
	ch <- 1
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan bool)

	go sendbool(ch2)
	go sendint(ch1)

	select {
	case i := <-ch1:
		fmt.Println(i)

	case b := <-ch2:
		fmt.Println(b)

	default:
		fmt.Println("Bye")
	}
}
