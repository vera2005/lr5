package main

import (
	"bufio"
	"fmt"
	"os"
)

func removeDuplicates(inputStream chan string, outputStream chan string) {
	prev := ""
	for v := range inputStream {
		if prev == "" || v != prev {
			outputStream <- v
		}
		prev = v
	}
	close(outputStream)
}

func main() {
	inputStream := make(chan string, 5)
	outputStream := make(chan string, 5)
	//для теста создаем каналы с типом строка
	// c размером буфера 5
	s := bufio.NewScanner(os.Stdin)
	//заполнение канала через ввод с клавиатуры
	fmt.Printf("Put %d strings\n", cap(inputStream))
	for i := 0; i < cap(inputStream); i++ {
		s.Scan()
		a := s.Text()
		inputStream <- a
	}
	close(inputStream) // закрываем канал ввода
	removeDuplicates(inputStream, outputStream)
	fmt.Printf("In outputStream: ")
	for v := range outputStream {
		fmt.Print(v, "; ")
	}

}
