package main

import "fmt"

func checkNumber(num int) {
	if num == 0 {
		fmt.Println("Ноль")
	} else if num > 0 {
		fmt.Println("Число положительное")
	} else {
		fmt.Println("Число отрицательное")
	}
}

func main() {
	var num int
	fmt.Println("Enter a number (positive or negative):")
	fmt.Scan(&num)

	checkNumber(num)
}
