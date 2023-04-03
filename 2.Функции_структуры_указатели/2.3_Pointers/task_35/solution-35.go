package main

import "fmt"

func test(x1 *int, x2 *int) {
	result := *x1 * *x2
	fmt.Println(result)
}

func main() {
	x := 3
	y := 6
	test(&x, &y)
}
