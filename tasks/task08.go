// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0

package main

import (
	"fmt"
)

// функция смены бита принимает число, позицию бита (i в условии) и параметр val (true - установить 1, false - установить 0)
func changeBit(number int64, pos int, val bool) int64 {
	if val {
		// нужно установить в 1, поэтому используем операцию побитового ИЛИ с маской вида 0000100, где только один бит равен 1 на изменяемой позиции
		return number | (1 << (pos - 1))
	} else {
		// нужно установить в 0, поэтому используем операцию побитового И, где маска имеет вид 1111011, где бит на изменяемой позиции равен 0
		return number & ^(1 << (pos - 1))
	}
}

func main() {
	var number int64
	number = 135

	fmt.Printf("Number %d in binary is %b\n", number, number)
	number = changeBit(number, 3, false)
	fmt.Printf("Number %d in binary is %b\n", number, number)
	number = changeBit(number, 3, true)
	fmt.Printf("Number %d in binary is %b\n", number, number)
	number = changeBit(number, 4, true)
	fmt.Printf("Number %d in binary is %b\n", number, number)
	number = changeBit(number, 1, true)
	fmt.Printf("Number %d in binary is %b\n", number, number)

}
