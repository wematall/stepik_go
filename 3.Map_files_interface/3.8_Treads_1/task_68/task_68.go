package main

// Внутри функции main (функцию объявлять не нужно),
// вам необходимо в отдельной горутине
//вызвать функцию work() и дождаться результатов ее
// выполнения.

// Функция work() ничего не принимает и не возвращает.

func main() {
	done := make(chan struct{})
	go func(d chan struct{}) {
		work()
		close(d)
	}(done)
	<-done
}

func work() {

}
