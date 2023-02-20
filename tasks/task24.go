// Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point
// с инкапсулированными параметрами x,y и конструктором

package main

import (
	"fmt"
	"math"
)

// структура Point с инкасулированным параметрами x, y
type Point struct {
	x float64
	y float64
}

// конструктор Point
func NewPoint(x float64, y float64) *Point {
	return &Point{x, y}
}

// функция нахождения дистанции между точками
func FindDistance(p1 *Point, p2 *Point) int {
	distance := math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
	return int(distance)
}

func main() {
	p1 := NewPoint(10.2, 24.3)
	p2 := NewPoint(2.1, 3.4)

	d := FindDistance(p1, p2)

	fmt.Println(d) // вывод 22 - правильный
}
