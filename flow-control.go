package main

import "fmt"

var count int = 0

func multipleReturn() (int, bool) {
	count++
	return count, true
}

func flowControl() {
	// Go has only four flow control statements
	if myCount, ok := multipleReturn(); ok {
		fmt.Println("Count is", myCount)
	}

	// traditional for with variable declaration
	for i := 0; i < 10; i++ {
		fmt.Println("i is", i)
	}

	// for like while!
	for true {
		if count > 10 {
			break
		}
		count++
	}

	mySlice := []string{"a", "b", "c", "d"}
	// iterating over a collection
	for index, value := range mySlice {
		fmt.Println("For index", index, "value is", value)
	}

	// switch
	myString := "z"
	switch myString {
	case "a":
		fmt.Println("It is A")
	case "b":
		fmt.Println("It is B")
	default:
		fmt.Println("Actually it is", myString)
	}

	// there is one more that will we see next!
}

func main() {
	flowControl()
}
