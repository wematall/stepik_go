package main

import "fmt"

func zero(x *int) {
	*x = 100
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x)
}
