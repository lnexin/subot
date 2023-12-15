package main

import (
	"fmt"
)

func main() {
	var a float64 = 1.0
	var b float64 = 1.00

	var c float64 = 2

	fmt.Println(a == b)
	fmt.Println((a + b) == c)
}
