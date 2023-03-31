package main

import "fmt"

func main() {
	var n, a int
	array := [100]int{}

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		array[i] = a
		if i%2 == 0 {
			fmt.Print(array[i], " ")
		}
	}

}
