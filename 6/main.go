package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	var c2 chan string = make(chan string) // same thing but more verbose

	go func() {
		for {
			c1 <- "Every 500 ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2 seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	// for {
	// 	fmt.Println(<-c1)
	// 	fmt.Println(<-c2)
	// }
	// This outputs
	// Every 500ms
	// Every 2 seconds
	// Every 500ms
	// Every 2 seconds
	// Every 500ms
	// Every 2 seconds
	// because they have to wait one after another. even though the 500ms one is faster
	// this is where select comes in
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
