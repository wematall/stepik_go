package main

import "fmt"

func main() {
	var n int
	prev_num := 0 // previous number of fibo
	curr_num := 1 // current number of fibo
	next_num := 0 // next number of fibo
	position := 0 // position of number in fibo row

	fmt.Scan(&n)

	for {
		next_num = prev_num + curr_num
		position += 1

		if curr_num == n {
			fmt.Println(position)
			break
		}
		if curr_num > n {
			fmt.Println("-1")
			break
		}

		prev_num = curr_num
		curr_num = next_num
	}
}
