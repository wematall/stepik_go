package main

import "fmt"

func main() {
	var n, el int
	var mySlice []int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&el)
		mySlice = append(mySlice, el)
	}

	fmt.Println(mySlice[3])
}
