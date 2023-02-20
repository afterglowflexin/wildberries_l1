// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func reverse(s string) string {
	// разделяем строку в слайс, разделитель это пробел
	sent := strings.Split(s, " ")

	var res string

	// идем по слайсу в обратном порядке и присоединяем слова к ответу
	for i := len(sent) - 1; i >= 0; i-- {
		res = res + " " + sent[i]
	}
	return res
}

func main() {
	s1 := "snow dog sun"
	s2 := "wb go l1"
	s3 := "phone call morning chair table"

	fmt.Println(reverse(s1))
	fmt.Println(reverse(s2))
	fmt.Println(reverse(s3))
}
