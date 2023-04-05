package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("ошибка")

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}
func main() {
	var a, b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	divide(a, b)
}
