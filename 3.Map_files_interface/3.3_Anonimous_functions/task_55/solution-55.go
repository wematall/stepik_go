package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x uint
	p := 100
	fmt.Scan(&x)

	fn := func(a uint) uint {
		var s string = strconv.FormatUint(uint64(a), 10)
		var z string
		for _, v := range s {
			i, _ := strconv.Atoi(string(v))
			if i != 0 && i%2 == 0 {
				q := strconv.Itoa(i)
				z += q
			}
		}
		g, err := strconv.Atoi(z)
		if err != nil {
			panic(err)
		}

		if g == 0 {
			return uint(p)
		}

		return uint(g)
	}

	fmt.Print(fn(x))
}
