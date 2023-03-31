package main

import "fmt"

func main() {
	var num int
	var num_amount = 0
	var biggest_num = 0

	for {
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		if biggest_num == num {
			num_amount++
		}
		if biggest_num < num {
			biggest_num = num
			num_amount = 1
		}
	}
	fmt.Println(num_amount)
}
