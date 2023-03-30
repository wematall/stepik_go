package main

import "fmt"

func main() {
	var a float64
	fmt.Scan(&a)

	if a > 10000 {
		fmt.Printf("%e \n", a)
	} else if a <= 0 {
		fmt.Printf("число %2.2f не подходит \n", a)
	} else {
		a = a * a
		fmt.Printf("%.4f \n", a)
	}
}
