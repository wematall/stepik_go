package main

import "fmt"

func main() {
	// first example
	sum := 0
	for i := 1; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)

	// second example
	var i = 1
	for ; i < 10; i++ {
		fmt.Println(i * i)
	}

}
