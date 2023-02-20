// Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
// которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров

package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const workersCount = 5

// функция воркера, читает данные из канала, имитирует работу, выводит данные в stdout
func worker(id int, ch chan int) {
	for {
		// воркер завершается, если канал закрыт. Еще можно было завершать работу воркеров аналогично с горутиной в main
		//  с помощью передачи контекста и оператора select
		x, ok := <-ch
		if !ok {
			fmt.Fprintf(os.Stdout, "Worker %v stopped\n", id)
			return
		}
		fmt.Fprintf(os.Stdout, "Worker %v started job %v\n", id, x)
		time.Sleep(2 * time.Second)
		fmt.Fprintf(os.Stdout, "Worker %v finished job %v\n", id, x)
	}

}

// функция запуска воркеров. Принимает на вход количество воркеров, которые надо запустить, и канал, из которого читать
func startWorkers(number int, ch chan int) {
	for i := 1; i <= number; i++ {
		go worker(i, ch)
	}
}

func main() {
	// создаем канал для передачи данных
	ch := make(chan int)

	// создаем контекст, который отмечается как Done при получении указанных сигналов
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	// запускаем постоянную запись данных в канал
	go func() {
		for {
			// с помощью оператора select горутина закроет канал и завершится в случае получения сигнала от операционной системы
			select {
			case <-ctx.Done():
				close(ch)
				fmt.Fprint(os.Stdout, "Channel closed\n")
				return
			default:
				ch <- rand.Int()
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// запускаем воркеров
	startWorkers(workersCount, ch)

	// прописываем graceful shutdown для того, чтобы worker'ы успели доделать свою работу
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan
	time.Sleep(2 * time.Second)
	fmt.Println("Received terminate, graceful shutdown.", sig)
}
