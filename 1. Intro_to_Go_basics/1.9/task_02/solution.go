// По данному трехзначному числу определите, все ли его цифры различны.
// На вход подается одно натуральное трехзначное число.
// Выведите "YES", если все цифры числа различны, в противном случае - "NO".

package main

import "fmt"

func compare(a, b, c int) {
	if a != b && b != c && c != a {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func main() {
	var num int
	var a, b, c int

	fmt.Println("Enter a 3 digit number (for example: 123):")
	fmt.Scan(&num)

	a = num / 100
	b = (num % 100) / 10
	c = num % 10

	compare(a, b, c)
}
