package main

import (
	"fmt"
	"time"
)

func runApp() {
	c := make(chan string)
	var x string = "test"
	go count(&x, c)

	msg := <-c
	fmt.Println(msg)
}

func main() {
	runApp()
}

func count(thing *string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- *thing
		time.Sleep(time.Millisecond * 500)
	}
}
