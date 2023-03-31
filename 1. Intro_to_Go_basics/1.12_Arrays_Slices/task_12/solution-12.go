package main

import "fmt"

func main() {
	array := [5]int{}
	var a, maxEl int
	// idxEl := 0

	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}

	maxEl = array[0]
	for _, el := range array {
		if maxEl < el {
			maxEl = el
			// idxEl = idx
		}
	}

	fmt.Println(maxEl)

}
