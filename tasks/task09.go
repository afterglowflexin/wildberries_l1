// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout

package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	// 2 канала из условия
	ch1 := make(chan int)
	ch2 := make(chan int)

	// массив чисел
	arr := []int{1, 2, 3, 4, 5, 6}

	// в горутине запускаем процесс записи чисел из массива в первый канал
	go func() {
		for _, x := range arr {
			ch1 <- x
		}
	}()

	// в следующей горутине данные берутся из первого канала и значения x*2 записываются во второй канал
	go func() {
		for x := range ch1 {
			ch2 <- x * 2
		}
	}()

	// в горутине Main данные читаются из второго канала и выводятся в stdout
	for x := range ch2 {
		fmt.Fprintf(os.Stdout, "%v\n", x)
	}

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)

	<-sigChan
}
