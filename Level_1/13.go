package main

// Поменять местами два числа без создания временной переменной
// Используем побитовое исключающее ИЛИ (XOR)

import (
	"fmt"
)

func swap2int(a, b *int) {
	fmt.Printf("До свапа a: %08b  b: %08b\n", *a, *b)
	*a = *a ^ *b //XOR-им первый раз
	fmt.Printf("После первого XOR a^b a: %08b  b: %08b\n", *a, *b)
	*b = *a ^ *b //XOR-им второй раз
	fmt.Printf("После второго XOR a^b a: %08b  b: %08b\n", *a, *b)
	*a = *a ^ *b //XOR-им третий раз
	fmt.Printf("После третьего XOR a^b a: %08b  b: %08b\n", *a, *b)
}

func main() {
	a, b := 3, 7      // Инициализируем переменные a и b значениями 3 и 7
	fmt.Println(a, b) // Выводим исходные значения a и b (3 7)

	swap2int(&a, &b) // Вызываем функцию swap2int, передавая адреса a и b

	fmt.Print(a, b) // Выводим значения a и b после обмена (7 3)
}
