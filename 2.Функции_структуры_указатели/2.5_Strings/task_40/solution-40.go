package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b string

	fmt.Scan(&a)
	fmt.Scan(&b)

	res := strings.Index(a, b)
	fmt.Println(res)
}
