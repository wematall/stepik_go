package main

import "fmt"

func test(a *int, b *int) {
	*a, *b = *b, *a
	fmt.Println(*a, *b)
}

func main() {
	x := 2
	y := 4
	test(&x, &y)
}
