package main

import "fmt"

func main() {
	var a, b, min, minsAmount int
	array := []int{}

	fmt.Scan(&a)

	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		array = append(array, b)
	}
	min = array[0]
	minsAmount = 0
	for i, el := range array {
		if min == array[i] {
			minsAmount += 1
		}

		if min > array[i] {
			minsAmount = 0
			min = el
			minsAmount += 1
		}

	}

	fmt.Println(minsAmount)
}
