package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Println("Enter first number a: ")
	fmt.Scan(&a)
	fmt.Println("Enter second number b: ")
	fmt.Scan(&b)

	sum := 0
	for i := a; i <= b; i++ {
		sum = sum + i
	}
	fmt.Println(sum)

}
