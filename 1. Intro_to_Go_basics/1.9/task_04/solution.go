package main

import "fmt"

func happyTicket(num int) {
	var d1, d2, d3, d4, d5, d6 int
	d1 = num / 100000
	d2 = (num % 100000) / 10000
	d3 = (num % 10000) / 1000
	d4 = (num % 1000) / 100
	d5 = (num % 100) / 10
	d6 = num % 10

	if d1+d2+d3 == (d4 + d5 + d6) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func main() {
	var num int

	fmt.Println("Enter a number (6 digits):")
	fmt.Scan(&num)

	// num = num % 10000
	// fmt.Println(num)
	happyTicket(num)

}
