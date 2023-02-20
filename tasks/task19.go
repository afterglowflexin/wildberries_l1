// Разработать программу, которая переворачивает подаваемую на ход строку
// например: «главрыба — абырвалг». Символы могут быть unicode.

package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(s string) string {
	// посчитаем кол-во рун в строке. Аналог len, но т.к. могут быть Unicode символы, считаем руны
	rcount := utf8.RuneCountInString(s)

	// создаем новый массив рун длинной rcount
	arr := make([]rune, rcount)

	// итерируемся по данной строке с помощью range (таким образом итерируемся по рунам) и заполняем arr с конца
	for _, x := range s {
		arr[rcount-1] = x
		rcount--
	}

	// возвращаем строковое выражение полученного массива рун
	return string(arr)
}

func main() {
	s1 := "главрыба"
	s2 := "crocodile"
	fmt.Println(reverse(s1)) // выводит "абырвалг"
	fmt.Println(reverse(s2)) // выводит "elidocorc"
}
