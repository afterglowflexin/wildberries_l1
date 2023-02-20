// Реализовать все возможные способы остановки выполнения горутины

// Способы:
// 1. Через контекст ctx.Done()
// 2. Горутины слушают канал, проверяют, если он закрылся, то останавливают свою работу
// 3. Завершение main горутины

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// первый способ, когда в горутину передается контекст, и при его завершении горутины останавливают свою работу
func doSmth1(ctx context.Context, id int) {

	for {
		select {
		// при выполнении контекста горутина завершится
		case <-ctx.Done():
			fmt.Printf("Goroutine 1.%v stopped its work\n", id)
			return
		default:
			fmt.Printf("Goroutine 1.%v doing work\n", id)
			time.Sleep(1 * time.Second)
			fmt.Printf("Goroutine 1.%v finished work\n", id)
		}
	}

}

// второй способ, когда горутина слушает канал и при его закрытии останавливает свою работу
// можно сделать отдельный канал специально для закрытия.
func doSmth2(ch chan int, id int) {

	for {
		// если канал закрывается, вторым значением при чтении из канала горутина получит false и завершит работу
		x, ok := <-ch
		if !ok {
			fmt.Printf("Goroutine 2.%v stopped listening\n", id)
			return
		}
		fmt.Printf("Goroutine 2.%v doing work with %v\n", id, x)
		time.Sleep(1 * time.Second)
		fmt.Printf("Goroutine 2.%v finished work with %v\n", id, x)
	}
}

func main() {
	// контекст для первого способа
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	// канал для второго способа
	ch := make(chan int)

	// запускаются три горутины с завершением с использованием контекста
	for i := 1; i <= 3; i++ {
		go doSmth1(ctx, i)
	}

	// запускается пять горутин с завершением с использованием закрытия канала
	for i := 1; i <= 5; i++ {
		go doSmth2(ch, i)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	<-sigChan
	// при получении сигнала от системы канал закрывается
	close(ch)
	// время для завершения работы горутин
	time.Sleep(1 * time.Second)
}
