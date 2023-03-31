package main

import "fmt"

func main() {
	var a, n int
	sum := 0

	for fmt.Scan(&n); n > 0; n = n / 10 {
		sum = sum + (n % 10)
	}
	fmt.Println(sum)
	if sum > 9 {
		a = sum
		sum = 0
		for i := a; i != 0; i = i / 10 {
			sum = sum + (i % 10)
		}
	}
	fmt.Println(sum)

}
