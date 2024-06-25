package main

//Реализовать бинарный поиск встроенными методами языка

import "fmt"

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2 // находим индекс среднего элемента

		if arr[mid] == target {
			return mid // элемент найден, возвращаем его индекс
		} else if arr[mid] < target {
			left = mid + 1 // искать в правой половине
		} else {
			right = mid - 1 // искать в левой половине
		}
	}

	return -1 // элемент не найден
}

func main() {
	sortedArray := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	target := 23

	index := binarySearch(sortedArray, target)

	if index != -1 {
		fmt.Printf("Элемент %d найден по индексу %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден в массиве\n", target)
	}
}
