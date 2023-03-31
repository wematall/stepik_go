package main

import "fmt"

func main() {
	var n, a int
	evenNumbers := 0

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if a > 0 {
			evenNumbers += 1
		}
	}
	fmt.Println(evenNumbers)
}
