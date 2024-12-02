package main

import (
	"fmt"
	"time"
)

func doesntRunBothCalls() {
	x := "sheep"
	count(&x)
	y := "other"
	count(&y)
}

func runsBothUsingGoroutines() {
	x := "sheep"
	go count(&x) // just add go in front and it will
	// split up the program into 2. main and the goroutine we just made
	y := "other"
	count(&y)
}

func main() {
	// doesntRunBothCalls()
	runsBothUsingGoroutines()
}

func count(x *string) {
	for i := 1; true; i++ {
		fmt.Println(i, *x)
		time.Sleep(time.Millisecond * 500)
	}
}
