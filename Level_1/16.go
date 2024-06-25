package main

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

import "fmt"

func quicksort(arr []int) []int { //алгоритм сортировки
	if len(arr) <= 1 { //проверка на пустой массив или на массив с 1 элеменом
		return arr
	}

	stack := [][]int{{0, len(arr) - 1}} //тут храним подмассивы, каждый элемент стека это слайс из двух элементов: начальный и конечный индексы подмассива.

	for len(stack) > 0 { //работает пока есть не обработанные пожмассивы

		last := len(stack) - 1
		start, end := stack[last][0], stack[last][1] // извлекаем последний элемент из стека (информация о следующем подмассиве).

		stack = stack[:last] // удаляем обработанный подмассив из стека

		pivotIndex := partition(arr, start, end) //разделяем подмассив относительно опорного пункта, pivotIndex - индекс опорного элемента после разделения

		if pivotIndex-1 > start { // добавляем подмассивы в стек, если их размер больше 1
			stack = append(stack, []int{start, pivotIndex - 1}) // левый подмассив (слева от pivot).
		}
		if pivotIndex+1 < end {
			stack = append(stack, []int{pivotIndex + 1, end}) // правый подмассив (справа от pivot).
		}
	}

	return arr // возвращаем отсортированный массив.
}

func partition(arr []int, start, end int) int { // partition разделяет подмассив arr[start:end] относительно опорного элемента (pivot)
	pivot := arr[end] // выбираем последний элемент подмассива как опорный
	i := start - 1    // i - индекс последнего элемента, меньшего pivot

	for j := start; j < end; j++ { // проходим по всем элементам подмассива (кроме pivot)
		if arr[j] <= pivot { // если текущий элемент меньше или равен pivot
			i++ //увеличиваем i и меняем местами arr[i] и arr[j],чтобы элементы, меньшие pivot, были слева от него.
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[end] = arr[end], arr[i+1] // меняем местами pivot (arr[end]) и элемент на позиции i+1

	return i + 1 // возвращаем индекс pivot
}

func main() {
	unsorted := []int{5, 2, 9, 3, 6, 1, 8, 4, 7}
	sorted := quicksort(unsorted)
	fmt.Println(sorted) // Вывод: [1 2 3 4 5 6 7 8 9]
}
