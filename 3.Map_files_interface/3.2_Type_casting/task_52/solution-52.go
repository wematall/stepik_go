package main

import (
	"fmt"
)

// func tion to cast type
func convert(a int64) uint16 {
	return uint16(a)
}

func main() {
	var a int64
	fmt.Scan(&a)

	res := convert(a)
	fmt.Printf("a : %d, %T \n", a, a)
	fmt.Printf("result: %d, %T \n", res, res)
}
