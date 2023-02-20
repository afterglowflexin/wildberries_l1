//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования)

package main

import (
	"fmt"
	"math"
)

// родительская структура human
type human struct {
	height float64
	weight float64
	age    int
}

// метод родительской структуры human для подсчета индекса массы тела
func (h *human) countBodyMassIndex() {
	fmt.Println(h.weight / math.Pow(h.height, 2))
}

// дочерняя структура action встраивает методы human, внедряя human без указания названия поля
type action struct {
	human
}

func main() {
	// объявление human
	ilya := human{height: 1.85, weight: 73.1, age: 25}

	// объявление action
	action := action{human: ilya}

	// видим, что action может напрямую использовать методы human
	action.countBodyMassIndex()
}
