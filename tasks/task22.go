// Разработать программу, которая перемножает, делит, складывает,
// вычитает две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Инициализируем большие числа
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("14721749120749210741297214791", 10)
	b.SetString("389213821983921839218934523", 10)

	// выполняем операции умножения, деления, сложения и вычитания
	mul := new(big.Int).Mul(a, b)
	div := new(big.Int).Div(a, b)
	sum := new(big.Int).Add(a, b)
	sub := new(big.Int).Sub(a, b)

	fmt.Printf("Mul: %v\nDiv: %v\nSum: %v\nSub: %v", mul, div, sum, sub)
}
