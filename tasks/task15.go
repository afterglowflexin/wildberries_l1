// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
// Приведите корректный пример реализации.

package main

import "fmt"

var justString string

func someFunc() {
	v := "djlsajodjsaopjdopasjopdjasopjdopasjdopjsaopjdopasjodpanvklnaiosbvnioqbiobwpiqk;bvklabklbsiobqbvklqbioqbiowbqibvk"
	justString = v[:20]
}

func main() {
	someFunc()

	fmt.Println(justString)
}
