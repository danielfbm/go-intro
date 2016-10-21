package main

import (
    "fmt"
)

func Variables() {

    var i int = 0

    var j, k int = 1, 2

    z, x, c := 3, 4, 5

    fmt.Println(i)
    fmt.Println(j)
    fmt.Println(k)
    fmt.Println(z, x, c)

    var one string 

    fmt.Println(one)

    one = "one"

    var two, three string = "two", "three"

    four, five, six := "four", "five", "six"
    
    fmt.Println(one)
    fmt.Println(two)
    fmt.Println(three)
    fmt.Println(four, five, six)
}

func main() {

    Variables()
}