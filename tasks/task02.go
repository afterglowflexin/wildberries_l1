// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout

package main

import (
	"fmt"
	"os"
)

func main() {
	// объявление массива
	numbers := []int{2, 4, 6, 8, 10}

	// создание канала
	ch := make(chan int)

	// для каждого элемента массива вызывается горутина и квадрат числа записывается в канал
	for _, x := range numbers {
		go func(a int) {
			ch <- a * a
		}(x)
	}

	// в цикле выводим данные из канала; при каждом чтении из канала горутина main блокируется,
	// поэтому мы точно получим квадраты всех чисел перед ее завершением
	for i := 0; i < len(numbers); i++ {
		fmt.Fprintln(os.Stdout, <-ch)
	}
}
