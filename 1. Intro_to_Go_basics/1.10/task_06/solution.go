package main

import "fmt"

func main() {
	var num int

	for {
		fmt.Scan(&num)
		if num < 10 {
			continue
		}
		if num > 100 {
			break
		} else {
			fmt.Println(num)
		}
	}
}
