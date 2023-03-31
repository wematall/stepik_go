package main

import "fmt"

func main() {
	var sec, h, m int // seconds, hours, minutes
	fmt.Scan(&sec)

	h = sec / 3600
	m = (sec - (sec/3600)*3600) / 60
	fmt.Printf("It is %d hours %d minutes. \n", h, m)
}
