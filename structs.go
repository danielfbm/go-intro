package main

import "fmt"

type MyStruct struct {
	Name  string
	Count int
}

func structs() {

	var instance MyStruct = MyStruct{
		Name:  "John Doe",
		Count: 1,
	}

	pointer := &MyStruct{
		Name:  "John 'pointer' Doe",
		Count: 1,
	}

	Increment(instance)
	IncrementP(pointer)

	fmt.Println("Outside increment---")
	fmt.Println(instance)
	fmt.Println(pointer)

	instance.Increment()
	pointer.IncrementP()

	fmt.Println("Instance method---")
	fmt.Println(instance)
	fmt.Println(pointer)

	instance.IncrementP()
	fmt.Println("Instance method 2---")
	fmt.Println(instance)

	PrintSomeIncrementReturn(instance)
}

func Increment(inst MyStruct) {
	inst.Count++
	fmt.Println("Inside increment")
	fmt.Println(inst)
}

func IncrementP(inst *MyStruct) {
	inst.Count++
	fmt.Println("Inside increment")
	fmt.Println(inst)
}

func PrintSomeIncrementReturn(instance Incrementer) {
	fmt.Println(instance.Increment())
}

func main() {
	structs()
}

func (my MyStruct) Increment() int {
	my.Count++
	return my.Count
}

func (my *MyStruct) IncrementP() {
	my.Count++
}

type Incrementer interface {
	Increment() int
}
