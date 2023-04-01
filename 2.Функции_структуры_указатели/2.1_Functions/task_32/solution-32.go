package main

import "fmt"

func vote(x, y, z int) int {
	array := [3]int{x, y, z}
	count_0, count_1 := 0, 0

	for i := 0; i < len(array); i++ {
		if array[i] == 0 {
			count_0 += 1
		} else {
			count_1 += 1
		}
	}

	if count_0 > count_1 {
		return 0
	} else {
		return 1
	}
}

func main() {
	var a, b, c = 0, 0, 1

	result := vote(a, b, c)
	fmt.Println(result)
}
