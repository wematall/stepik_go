package main

import (
	"fmt"
	"math"
)

func hipotenuse(a, b float64) float64 {
	result := math.Sqrt(a*a + b*b)
	return result
}

func main() {
	var a, b float64

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println(hipotenuse(a, b))
}
