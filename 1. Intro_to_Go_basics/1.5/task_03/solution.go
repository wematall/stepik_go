package main

import "fmt"

func Check(number int) {
	if number < 0 || number > 360 {
		fmt.Println("Wrong number")
	}
}

// degrees to hours and minutes
// 30 degrees - 1 hour
// 30 degrees - 60 minutes

func main() {
	var num int // user number degrees
	var h int   // hours
	var m int   // minutes
	fmt.Println("Enter a number (0 < number < 360): ")
	fmt.Scan(&num)
	Check(num)

	h = num / 30
	m = (num - h*30) * 2

	fmt.Println("It is", h, "hours", m, "minutes.")
}
