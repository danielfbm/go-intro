package main

import "fmt"

/*
type error interface {
    Error() string
}
*/

type MyCustomType struct {
	ID    int
	Name  string
	Items []struct {
		ItemName string
		Price    float64
	}
}

func (my MyCustomType) Error() string {
	return fmt.Sprintf("MyCustomTypeError: %d - Name %v", my.ID, my.Name)
}

func PrintError(err error) {
	fmt.Println(err.Error())
}

func main() {
	instance := MyCustomType{
		ID:   1,
		Name: "MyName",
	}

	PrintError(instance)
}
