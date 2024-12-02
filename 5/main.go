package main

import (
	"fmt"
	"time"
)

func runApp() {
	c := make(chan string)
	var x string = "test"
	go count(&x, c)

	// for {
	// 	msg, open := <-c
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }
	for msg := range c {
		fmt.Println(msg)
	}
}

func main() {
	runApp()
}

func count(thing *string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- *thing
		time.Sleep(time.Millisecond * 500)
	}
	close(c) // never close the channel unless you are sender
	// the reciever cant know if the sender is done sending
	// so closing c here is good practice
}
