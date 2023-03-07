package main

import "fmt"

func printFirstDigit(num int) {
	var a int

	if num == 10000 {
		a = num / 10000
	} else if num <= 9999 && num >= 1000 {
		a = num / 1000
	} else if num <= 999 && num >= 100 {
		a = num / 100
	} else if num <= 99 && num >= 10 {
		a = num / 10
	} else {
		a = num
	}
	fmt.Println(a)
}

func main() {
	var num int

	fmt.Println("Enter a positive number (10000 or less):")
	fmt.Scan(&num)

	printFirstDigit(num)
}
