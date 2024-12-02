package main

import (
	"fmt"
	"sync"
	"time"
)

func runApp() {
	var wg sync.WaitGroup
	wg.Add(1)
	var x string = "test"
	go func() {
		count(&x)
		wg.Done()
	}()
	wg.Wait()
}

func main() {
	runApp()
}

func count(x *string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, *x)
		time.Sleep(time.Millisecond * 500)
	}
}
