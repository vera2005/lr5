package main

import (
	"fmt"
)

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	ans := make(chan int) //создание возвращаемого канала

	// чтобы сделать функцию неблокирующей, убираем
	// в  анонимную функцию, вызываемую через горутину
	// логику select т.к. в противном случае выполнение функции блокируется до момента поступления значения
	go func() {
		defer close(ans) //
		select {
		case a := <-firstChan:
			ans <- a * a
		case a := <-secondChan:
			ans <- a * 3
		case <-stopChan:
			return
		}
	}()
	return ans
}

func main() {
	first := make(chan int)
	second := make(chan int)
	stop := make(chan struct{})
	var a, b int
	fmt.Println("Put 1 for square, 2 for multiplication by 3 or 3 for stop:")
	fmt.Scan(&a)
	if a == 1 || a == 2 {
		fmt.Println("Put a number")
		fmt.Scan(&b)
	}
	var ansChan <-chan int
	switch a {
	case 1:
		ansChan = calculator(first, second, stop)
		first <- b

	case 2:
		ansChan = calculator(first, second, stop)
		second <- b

	case 3:
		close(stop)                               // Закрываем канал и сигнализируем остановку
		ansChan = calculator(first, second, stop) // Запускаем функцию для остановки
	}
	ans, ok := <-ansChan // Получаем результат и статус
	if ok {
		fmt.Println("Result is: ", ans)
	} else {
		fmt.Println("No result, calculator was stopped.")
	}

}
