// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество

package main

import "fmt"

func group() {

}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	// используется пустая структура, тк она ничего не весит
	res := make(map[string]struct{})

	for _, x := range arr {
		res[x] = struct{}{}
	}

	for x := range res {
		fmt.Printf("%v ", x)
	}
}
