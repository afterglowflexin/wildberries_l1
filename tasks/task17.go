// Реализовать бинарный поиск встроенными методами языка

package main

import "fmt"

// стандартная функция бинарного поиска с двумя указателями на концы рассматриваемой части массива
func binSearch(arr []int, val int) int {
	// l (left) - указатель на левый конец, r (right) - указатель на правый конец
	l := 0
	r := len(arr) - 1

	// в цикле считаем середину между l и r и сравниваем искомое значение и значение на позиции mid
	// если arr[mid] больше val, значит искомое значение может находится в левой половине рассматриваемой части массива, есди меньше - в правой
	for l <= r {
		mid := (l + r) / 2

		if arr[mid] == val {
			return mid
		} else if arr[mid] < val {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	// возвращаем -1, если искомого числа нет в данном массиве
	return -1
}

func main() {
	// объявляем сортированный массив
	arr := []int{-42, -30, -23, -17, -3, 0, 1, 4, 12, 16, 26, 31}

	fmt.Printf("Actual is %d, expected is 0\n", binSearch(arr, -42))
	fmt.Printf("Actual is %d, expected is 1\n", binSearch(arr, -30))
	fmt.Printf("Actual is %d, expected is 5\n", binSearch(arr, 0))
	fmt.Printf("Actual is %d, expected is 8\n", binSearch(arr, 12))
	fmt.Printf("Actual is %d, expected is 11\n", binSearch(arr, 31))
	fmt.Printf("Actual is %d, expected is -1\n", binSearch(arr, -54))
	fmt.Printf("Actual is %d, expected is -1\n", binSearch(arr, 61))
}
