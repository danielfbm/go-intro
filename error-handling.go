package main

import (
	"fmt"
	"strconv"
)

func main() {
	myText := "1"

	var value int64
	var err error
	if value, err = strconv.ParseInt(myText, 10, 0); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("value is", value)
	}

	myText = "s"

	if value, err = strconv.ParseInt(myText, 10, 0); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("value is", value)
	}

	value = 10
	panicAttack(value)

}

func panicAttack(val int64) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("oh goodness, just recovered...")
			fmt.Println(r)
		}
	}()
	if val == 0 {
		panic("Oh my God, it is zero")
	} else {
		fmt.Println("Got", val)

	}
	panicAttack(val - 1)

}
