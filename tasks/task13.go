// Поменять местами два числа без создания временной переменной

package main

import "fmt"

func main() {
	first := 15
	second := 20

	fmt.Printf("First is %v, second is %v\n", first, second) // Вывод "First is 15, second is 20"

	first, second = second, first

	fmt.Printf("First is %v, second is %v\n", first, second) // Вывод "First is 20, second is 15"
}
