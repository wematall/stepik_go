package main

import "sync"

// Внутри функции main
// вам необходимо в отдельных горутинах
// вызвать функцию work() 10 раз
//и дождаться результатов выполнения вызванных
// функций.
// Функция work() ничего не принимает и не возвращает.

func main() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		// Запускаем 10 экземпляров горутины, увеличивающей счетчик на 1
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			work()
		}(wg)
	}

	wg.Wait()
}

func work() {}
