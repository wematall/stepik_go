package main

import "fmt"

func fibonacci(n int) int {
	prev_num := 0 // previous number of fibo
	curr_num := 1 // current number of fibo
	next_num := 0 // next number of fibo
	position := 1 // position of number in fibo row

	for position != n {
		position += 1
		next_num = prev_num + curr_num

		prev_num = curr_num
		curr_num = next_num
	}
	return curr_num
}

func main() {
	var n int

	fmt.Scan(&n)

	result := fibonacci(n)
	fmt.Println(result)
}
