package main

import "fmt"

func main() {
	var a, b int
	zeroNumAmount := 0

	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		fmt.Scan(&b)

		if b == 0 {
			zeroNumAmount += 1
		}
	}
	fmt.Println(zeroNumAmount)
}
