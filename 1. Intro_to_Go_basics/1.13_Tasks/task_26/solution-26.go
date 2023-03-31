package main

import "fmt"

func main() {
	var n int
	num := 1
	fmt.Scan(&n)

	for {
		fmt.Print(num, " ")
		num *= 2
		if num >= n {
			break
		}
	}
}
