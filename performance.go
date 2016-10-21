package main

import (
	"fmt"
	"time"
)

func fibonacci(number int) int {
	if number > 1 {
		return fibonacci(number-1) + fibonacci(number-2)
	}
	return number
}

func main() {
	// var number int
	start := time.Now()
	for i := 0; i < 40; i++ {
		fmt.Println(i, fibonacci(i))
	}
	fmt.Println(time.Now().Sub(start).Seconds())
}
