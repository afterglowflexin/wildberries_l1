// Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"sync"
)

// 1-ый способ с кастомной мапой, Mutex для организации конкурентной записи
type custom_map struct {
	values map[int]bool
	sync.Mutex
}

// конструктор
func newMap() *custom_map {
	return &custom_map{values: make(map[int]bool)}
}

func main() {
	m := newMap()
	wg := sync.WaitGroup{}

	// Запуск 10 горутин, которые конкурентно записывают данные в map
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(m *custom_map, i int) {
			defer wg.Done()
			m.Lock()
			defer m.Unlock()

			m.values[i] = true
		}(m, i)
	}

	wg.Wait()

	fmt.Println(m.values) // вывод map[0:true 1:true 2:true 3:true 4:true 5:true 6:true 7:true 8:true 9:true]

	// 2-ой способ с встроенной sync.Map
	m2 := &sync.Map{}
	wg2 := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg2.Add(1)

		go func(m *sync.Map, i int) {
			// подсчет квадратов чисел
			m.Store(i, i*i)
			wg2.Done()
		}(m2, i)
	}

	wg2.Wait()

	m2.Range(func(key any, value any) bool {
		fmt.Printf("%v:%v ", key, value)
		return true
	}) // вывод 9:81 8:64 6:36 7:49 4:16 0:0 3:9 2:4 1:1 5:25
}
