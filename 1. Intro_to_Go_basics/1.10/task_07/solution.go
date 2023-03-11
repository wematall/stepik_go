package main

import "fmt"

func main() {
	var x, p, y int
	fmt.Scan(&x, &p, &y)

	i := 0
	for {
		if x < y {
			x = x + (x*p)/100 // считаем увеличение вклада на p процентов
		} else {
			break
		}
		i++
	}
	fmt.Println(i)
}
