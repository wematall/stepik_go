package main

import "fmt"

func main() {
	var a int

	fmt.Println("Enter a number:")
	fmt.Scan(&a)

	result := (a % 100) / 10
	fmt.Println(result)
}
