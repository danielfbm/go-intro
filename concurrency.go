package main

import "fmt"

func concurrency() {

	// lets first create a channel with a buffer
	tasks := make(chan string, 20)

	// create another one to receive the results
	results := make(chan string, 20)

	workers := []int{1, 2, 3, 4}

	// inserting tasks inside the channel
	for task := 0; task < 10; task++ {
		tasks <- fmt.Sprintf("Task %d", task)
	}

	for _, w := range workers {
		// starting one goroutine for each worker
		go work(w, tasks, results)
	}
	//closing this channel so that workers can quit
	close(tasks)

	//lets print the results
	fmt.Println("Will print the results")

	// collecting the results
	for res := 0; res < 10; res++ {
		fmt.Println("Result:", <-results)
	}
}

func work(workerID int, tasks chan string, results chan string) {
	// worker will block until a new task arrives in the channel
	for t := range tasks {
		// simple task as example
		results <- fmt.Sprintf("Worker %d got %v", workerID, t)
	}
}

func main() {
	concurrency()
}
