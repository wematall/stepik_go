package main

import "fmt"

func minimumFromFour(a, b, c, d int) int {
	if a < b && a < c && a < d {
		return a
	} else if b < a && b < c && b < d {
		return b
	} else if c < a && c < b && c < d {
		return c
	} else {
		return d
	}
}

func main() {
	a, b, c, d := 4, 5, 6, 7
	result := minimumFromFour(a, b, c, d)
	fmt.Println(result)
}
