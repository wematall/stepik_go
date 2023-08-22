package main

// Напишите функцию которая принимает канал и число
// N типа int.
// Необходимо вернуть значение N+1 в канал.
// Функция должна называться task().

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go task(c, 5)
	time.Sleep(time.Second * 1)
	fmt.Println(<-c)
}

func task(c chan int, N int) {
	c <- N + 1
}
