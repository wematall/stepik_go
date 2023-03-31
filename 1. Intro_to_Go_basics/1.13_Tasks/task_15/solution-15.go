package main

import "fmt"

func main() {
	var n, sum int

	fmt.Scan(&n)

	sum = n / 100
	// fmt.Println(sum)
	sum = sum + (n-(n/100)*100)/10
	// fmt.Println(sum)
	sum = sum + (n % 10)

	fmt.Println(sum)
}
