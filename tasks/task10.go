// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

package main

import "fmt"

// функция принимает массив чисел и возвращает мапу как в примере
//
// сделал, как было указано в примере, что значения, допустим, от -30 до -20 принадлежать к подмножеству -20,
// но при таком подсчете в множестве 0 будут числа от -10 до 10, т.е. шаг 10 градусов не соблюдается
// Поэтому есть еще закомментированный вариант, где сохраняется шаг.

func degrees(a []float32) map[int][]float32 {
	m := make(map[int][]float32)

	// для каждого числа в массиве считается индекс его множества с помощью приведения float32 к int и подсчета цело
	for _, x := range a {
		// var index int
		// if x < 0 {
		// 	index = int(x-10) / 10 * 10
		// } else {
		// 	index = int(x) / 10 * 10
		// }

		index := int(x) / 10 * 10

		m[index] = append(m[index], x)
	}

	return m
}

func main() {
	arr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	fmt.Println(degrees(arr)) // вывод map[-20:[-25.4 -27 -21] 10:[13 19 15.5] 20:[24.5] 30:[32.5]]
}
