package main

import "fmt"

func main() {
	var a, b int
	max, count := 0, 0

	fmt.Scan(&a, &b)

	for i := b; i >= a; i-- {
		if i%7 == 0 {
			max = i
			count = 1
			fmt.Println(max)
			break
		}
	}
	if count == 0 {
		fmt.Println("NO")
	}
}
