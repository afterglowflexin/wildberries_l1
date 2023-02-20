// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

package main

import (
	"fmt"
	"strings"
)

func uniqueSymbols(s string) bool {
	// приводим все символы в строке к нижнему регистру, чтобы функция была регистронезависимой
	s = strings.ToLower(s)

	// создаем мапу, в которой будем хранить true для символов, которые уже встречались при обходе
	m := make(map[rune]bool)

	// проходим по всем символам в строке. Если символ уже встречался, возвращаем false.
	for _, x := range s {
		if m[x] {
			return false
		}
		m[x] = true
	}

	// если в цикле все символы были уникальными, значит, вовзращаем true
	return true
}

func main() {
	s1, s2, s3, s4, s5 := "abcd", "abCdefAaf", "aabcd", "aA", "AbCd" // должно вывести true, false, false, false, true

	fmt.Println(
		uniqueSymbols(s1),
		uniqueSymbols(s2),
		uniqueSymbols(s3),
		uniqueSymbols(s4),
		uniqueSymbols(s5),
	) // вывод true false false false true
}
