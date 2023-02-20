// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// задаем значение N
const n = 10

func main() {
	timeout := n * time.Second

	// задаем контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// создаем канал для передачи данных
	ch := make(chan int)

	// в отдельной горутине отправляем данные в канал каждые 2 секунды, передавая ей контекст
	go func(ctx context.Context) {
		for {
			// с помощью оператора select пишем данные и ждем завершение контекста
			select {
			case ch <- rand.Int():
				time.Sleep(2 * time.Second)
			// по истечении контекста завершаем цикл, закрываем канал и горутина завершается
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				close(ch)
				return
			}
		}
	}(ctx)

	// с другой стороны канала читаем данные и выводим их
	for x := range ch {
		fmt.Println(x)
	}
}
