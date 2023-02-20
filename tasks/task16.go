// Реализовать быструю сортировку массива (quicksort) встроенными методами языка

package main

import "fmt"

// функция quicksort принимает массив и индексы, которые являются концами рассматриваемой части массива
func quicksort(a []int, minIndex int, maxIndex int) []int {
	if minIndex >= maxIndex {
		return a
	}

	// найдем место, где в итоге будет стоять pivot
	pivotIndex := getPivotIndex(a, minIndex, maxIndex)

	// рекурсивно вызываем для двух подмассивов, которые разделяются pivot'ом
	quicksort(a, 0, pivotIndex-1)
	quicksort(a, pivotIndex+1, maxIndex)

	return a
}

// функция для нахождения пивота и сортировки относительно него
func getPivotIndex(a []int, minIndex int, maxIndex int) int {
	pivot := minIndex - 1

	// если есть хоть одно число меньше pivot (за pivot берется последнее число рассматриваемого массива - a[maxIndex]),
	// то позиция pivot смещается на 1
	for i := minIndex; i <= maxIndex; i++ {
		if a[i] < a[maxIndex] {
			// также переставляем число меньшее значения pivot a[maxIndex] перед числами большими a[maxIndex]
			pivot++
			swap(a, pivot, i)
		}
	}

	// ставим pivot на свое место
	pivot++
	swap(a, pivot, maxIndex)

	// получаем отсортированный относительно pivot'а массив
	return pivot
}

// функция для перестановки элементов массива
func swap(a []int, i int, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	a := []int{2, 8, 1, 0, -23, -34, -35, -100, 100, 4}
	fmt.Println(a)

	quicksort(a, 0, len(a)-1)
	fmt.Println(a) // выводит [-100 -35 -34 -23 0 1 2 4 8 100]
}
