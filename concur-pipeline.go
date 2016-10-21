package main

import "fmt"

type Pipe interface {
	Process(in chan int) chan int
}

type Sq struct{}

func (Sq) Process(in chan int) chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			out <- i * i
		}
		close(out)
	}()
	return out
}

type Add struct{}

func (Add) Process(in chan int) chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			out <- i + i
		}
		close(out)
	}()
	return out
}

type Pipeline struct {
	head chan int
	tail chan int
}

func (p *Pipeline) Enqueue(item int) {
	p.head <- item
}

func (p *Pipeline) Dequeue(handler func(int)) {
	for i := range p.tail {
		handler(i)
	}
}

func (p *Pipeline) Close() {
	close(p.head)
}

func NewPipeline(pipes ...Pipe) Pipeline {
	head := make(chan int)
	var next_chan chan int
	for _, pipe := range pipes {
		if next_chan == nil {
			next_chan = pipe.Process(head)
		} else {
			next_chan = pipe.Process(next_chan)
		}
	}
	return Pipeline{head: head, tail: next_chan}
}

func main() {
	pipeline := NewPipeline(Sq{}, Add{})
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Sending:", i)
			pipeline.Enqueue(i)
		}
		fmt.Println("Closing Pipeline.")
		pipeline.Close()
	}()
	pipeline.Dequeue(func(i int) {
		fmt.Println("Received:", i)
	})
}
