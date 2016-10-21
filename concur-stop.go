package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Worker our brave minion
type Worker struct {
	ID       int
	counter  int
	stopSign <-chan bool
}

// DoStuff that is where the magic happens
func (w *Worker) DoStuff() bool {
	// Select is the missing control statement, used to make decisions over
	// channel information
	select {
	case value := <-w.stopSign:
		fmt.Println("\nWorker", w.ID, "will stop now", value)
		return false
	default:
		w.counter++

		fmt.Printf("\nWorker %d - job %d", w.ID, w.counter)

	}
	return true
}

// WorkaholicMode starts to work non stop until receive a stop sign
func (w *Worker) WorkaholicMode() {
	for {
		if !w.DoStuff() {
			fmt.Println("Ok, time to go home")
			break
		}
		time.Sleep(time.Millisecond * time.Duration(500))
	}
}

func stopSignal() {

	stopChannel := make(chan bool)

	minions := []*Worker{
		&Worker{
			ID:       1,
			stopSign: stopChannel,
		},
		&Worker{
			ID:       2,
			stopSign: stopChannel,
		},
	}

	for _, w := range minions {
		go w.WorkaholicMode()
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Got it! will stop now")

		stopChannel <- true
		stopChannel <- true

		//
		go func() {
			time.Sleep(time.Second)
			os.Exit(0)
		}()
	})

	log.Fatal(http.ListenAndServe(":3030", nil))
}

func main() {
	stopSignal()
}
