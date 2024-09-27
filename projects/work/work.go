package main

import (
	"fmt"
	"sync"
	"time"
)

func work() {
	time.Sleep(time.Millisecond * 50)
	fmt.Println("done")
}

func main() {
	wait := new(sync.WaitGroup) // определяем группу горутин для совместного выполнения
	for i := 0; i < 10; i++ {
		wait.Add(1) // увеличиваем счетчик внутренних активных элементов
		// анонимная горутина
		go func(w *sync.WaitGroup) {
			defer w.Done() // отложенный сигнал о завршершении горутины
			work()
		}(wait)
	}
	wait.Wait() // ожидание выполнения всех горутин, блокировка main
	fmt.Println("Completed")
}
