package main

import "fmt"

func main() {
	// fmt.Println("Enter 10 numbers: ")
	var workArray [10]uint8
	var num, tmp, next uint8

	for i := 0; i < 10; i++ {
		fmt.Scan(&num)
		workArray[i] = num
	}

	for i := 0; i < 3; i++ {
		fmt.Scan(&num)
		tmp = workArray[num]
		fmt.Scan(&next)
		workArray[num] = workArray[next]
		workArray[next] = tmp
	}
	fmt.Println()
	for idx := range workArray {
		fmt.Print(workArray[idx], " ")
	}
}
