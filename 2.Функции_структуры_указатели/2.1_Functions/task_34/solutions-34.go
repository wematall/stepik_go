package main

import "fmt"

func sumInt(a ...int) (x, y int) {
	sum := 0
	count := 0

	for _, element := range a {
		sum += element
		count += 1
	}
	return count, sum
}

func main() {
	a, b := sumInt(1, 0)
	fmt.Println(a, b)
}
