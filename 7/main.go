package main

import "fmt"

func fibo(n int) int {
	if n <= 1 {
		return n
	}
	return fibo(n-1) + fibo(n-2)
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fibo(n)
	}
}

func main() {
	jobs := make(chan int, 30)
	results := make(chan int, 30)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i <= 45; i++ {
		jobs <- i
	}
	close(jobs) // we are in the sender
	// now we expect 20 items to appear
	for i := 0; i <= 45; i++ {
		fmt.Println(<-results)
	}
}
