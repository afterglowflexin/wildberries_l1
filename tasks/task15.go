// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
// Приведите корректный пример реализации.

package main

import "fmt"

var justString string

// В этом примере функция работает со строкой как с массивом байтов.
// Для корректного отображения символов, состоящих из более, чем одного байта, необходимо работать с рунами.
//
// Также в этой функции может произойти выход за границы массива
//
// Не следует работать в функции с глобальной переменной, лучше вовзвращать строку
func someFunc() {
	v := "ಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕ"
	justString = v[:10]
}

// корректный пример
func someFuncMod() string {
	v := []rune("ಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕಕ")

	var str string

	if len(v) >= 11 {
		str = string(v[:10])
	}

	return str
}

func main() {
	someFunc()
	str := someFuncMod()

	fmt.Printf("First variant: %v, second variant: %v", justString, str) // First variant: ಕಕಕ�, second variant: ಕಕಕಕಕಕಕಕಕಕ
}
