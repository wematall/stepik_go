package main

import "fmt"

func main() {
	var a, b int
	var average float64

	fmt.Scan(&a, &b)
	average = (float64)(a+b) / 2
	fmt.Println(average)
}
