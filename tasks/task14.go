// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel
// из переменной типа interface{}

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var (
		num int
		s   string
		b   bool
		ch  chan int
	)

	fmt.Println(reflect.TypeOf(interface{}(num))) // вывод int
	fmt.Println(reflect.TypeOf(interface{}(s)))   // вывод string
	fmt.Println(reflect.TypeOf(interface{}(b)))   // вывод bool
	fmt.Println(reflect.TypeOf(interface{}(ch)))  // вывод chan int
}
