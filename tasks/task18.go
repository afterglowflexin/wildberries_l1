// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
	"time"
)

// структура счетчика. Значение счетчика и mutex для работы в конкурентной среде
type counter struct {
	num int
	sync.Mutex
}

// конструктор счетчика
func newCounter() *counter {
	return &counter{num: 0}
}

// операция инкрементирования счетчика
func (c *counter) inc() {
	c.Lock()
	defer c.Unlock()

	c.num++
}

// функция, возвращающая значение счетчика
func (c *counter) val() int {
	return c.num
}

// функция с конкурентной работой
func doSmth(c *counter) {
	// создаем wait group для синхронизации горутин
	wg := sync.WaitGroup{}

	// выполняем 10 конкурентных операций, в которых инкрементируем счетчик
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			time.Sleep(1 * time.Second)
			c.inc()
			wg.Done()
		}()
	}

	wg.Wait()
}

func main() {
	// создаем новый счетчик
	counter := newCounter()

	// выполняем работу
	doSmth(counter)

	fmt.Println(counter.val()) // вывод 10
}
