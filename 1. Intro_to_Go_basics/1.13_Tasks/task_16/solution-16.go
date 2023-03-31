package main

import "fmt"

func main() {
	var a, b int // input and output numbers

	fmt.Scan(&a)

	b = (a % 10) * 100
	b = b + ((a-(a/100)*100)/10)*10
	b = b + (a / 100)

	fmt.Println(b)
}
