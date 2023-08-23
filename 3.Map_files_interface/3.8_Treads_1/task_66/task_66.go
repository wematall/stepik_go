package main

// Напишите функцию
// которая принимает канал и строку.
// Необходимо отправить эту же строку
// в заданный канал 5 раз, добавив к ней пробел.

// Функция должна называться task2().

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go task2(c, "apple")
	time.Sleep(time.Second * 1)
	fmt.Println(<-c)
}

func task2(c chan string, S string) {
	for i := 0; i < 5; i++ {
		c <- S + " "
	}
}
