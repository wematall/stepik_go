package main

import "fmt"

func scloniator(n int) string {
	if (n%10 == 0) || (n > 4 && n < 20) || (n%10 > 4 && n%10 < 10) {
		return "korov"
	} else if n == 1 || n%10 == 1 {
		return "korova"
	} else if (n > 1 && n < 5) || (n%10 > 1 && n%10 < 5) {
		return "korovy"
	} else {
		return "wrong korova"
	}
}

func main() {
	var n int
	var s string

	fmt.Scan(&n)
	s = scloniator(n)

	fmt.Printf("%d %s", n, s)
}
