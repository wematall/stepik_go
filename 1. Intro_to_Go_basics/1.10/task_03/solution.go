package main

import "fmt"

func main() {
	var sum = 0 // summ of two digits numbers
	var n int   // subsequence length
	var num int

	// define length of subsequence of numbers
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		if (num/10 > 0) && (num/100 == 0) && (num%8 == 0) {
			sum = sum + num
		}
	}
	fmt.Println(sum)
}
