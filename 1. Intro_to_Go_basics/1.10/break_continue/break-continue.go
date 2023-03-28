package main

import "fmt"

func main() {
	{
		var sum = 0
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				continue // переходим к следующей итерации
			}
			sum += i
			// sum = sum + i
		}
		fmt.Println("Summ: ", sum) // Сумма нечетных чисел
	}
	{
		var sum = 0
		for i := 1; i <= 9; i++ {
			if i > 4 {
				break // если число больше 4, выходим из цикла
			}
			sum += i
			// sum = sum + i
		}
		fmt.Println("Сумма: ", sum) // сумма чисел
	}
}
