// Реализовать собственную функцию sleep

package main

import (
	"context"
	"fmt"
	"time"
)

// функция sleep использует встроенную функцию time.After(), которая ждет истечения заданного времени
func sleep(d time.Duration) time.Time {
	return <-time.After(d)
}

// функция sleepCtx использует контекст с таймаутом и ждет его завершения
func sleepCtx(d time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), d)
	defer cancel()

	<-ctx.Done()
}

func main() {
	fmt.Println(time.Now())
	fmt.Println("Sleeping 3 seconds")

	t := sleep(3 * time.Second)
	fmt.Println(t)

	fmt.Println("Sleeping 3 seconds")

	sleepCtx(3 * time.Second)
	fmt.Println(time.Now())
}
