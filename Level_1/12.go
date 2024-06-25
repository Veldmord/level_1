package main

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество

import (
	"fmt"
)

func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}
	mapStr := make(map[string]bool)

	for _, word := range str {
		mapStr[word] = true
	}

	fmt.Println(mapStr)
}
