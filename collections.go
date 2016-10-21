package main

import "fmt"

func arrays() {

	// explicit type
	var first [3]string

	// inferred type with empty values
	auto := [3]string{}

	// automatically setting the size
	filled := [...]string{"a", "b", "c"}

	fmt.Println(first)
	fmt.Println(filled)
	fmt.Println(auto)

	// len() returns the length of the array
	fmt.Println(len(first))
	fmt.Println(len(auto))
	fmt.Println(len(filled))

	// cap() returns the capacity of the structure
	fmt.Println(cap(first))
	fmt.Println(cap(auto))
	fmt.Println(cap(filled))
}

func slices() {
	// Almost like array but with variable length and capacity

	newSlice := []string{}

	anotherSlice := make([]string, 0, 3)

	fmt.Println("newSlice ---")
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	fmt.Println("anotherSlice ---")
	fmt.Println(len(anotherSlice))
	fmt.Println(cap(anotherSlice))

	newSlice = append(newSlice, "a")
	newSlice = append(newSlice, "b")

	fmt.Println("newSlice ---")
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	anotherSlice = append(anotherSlice, "z")
	anotherSlice = append(anotherSlice, "x")
	anotherSlice = append(anotherSlice, "c")

	fmt.Println("anotherSlice --")
	fmt.Println(len(anotherSlice))
	fmt.Println(cap(anotherSlice))

	anotherSlice = append(anotherSlice, "v")

	fmt.Println("anotherSlice --")
	fmt.Println(len(anotherSlice))
	fmt.Println(cap(anotherSlice))

}

func main() {
	// arrays()
	slices()
}
