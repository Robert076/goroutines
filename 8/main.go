package main

import "fmt"

func isEven(x int) bool {
	return (x%2 == 0)
}

func worker(jobs <-chan int, results chan<- bool) {
	for i := range jobs {
		results <- isEven(i)
	}
	close(results)
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan bool, 100)

	go worker(jobs, results)

	for i := 0; i <= 100; i++ {
		jobs <- i
	}

	close(jobs)

	for i := range results {
		fmt.Println(i)
	}

}
