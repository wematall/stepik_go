package main

import "fmt"

func isALeapYear(year int) {
	var s string
	if (year%400 == 0) || (year%4 == 0 && year%100 != 0) {
		s = "YES"
	} else {
		s = "NO"
	}
	fmt.Println(s)
}

func main() {
	var year int
	fmt.Println("Enter year (up to 10000):")
	fmt.Scan(&year)
	isALeapYear(year)
}
