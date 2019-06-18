package main

import (
	"fmt"
	"math"
)

var n1 string = "mohamed "
var n2 float32 = 1.2

func init() {
	fmt.Println("First Code")
}
func main() {
	const s = 50
	fmt.Println(math.Pow(5, 2))
	a := 10
	defer fmt.Println(a)
	defer fmt.Println(n1)
	const (
		ad = iota
		bd
		cd
		sd
	)
	fmt.Println(bd, cd, sd)
}

// multipal return values
func add(a, b int) (int, float64) {
	return a + b, 1
}
