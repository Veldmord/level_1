package main

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

import (
	"fmt"
)

func sdvigBit(num int64, i int, b bool) int64 {
	mask := int64(1) << i //создается маска где установлен только i бит (вначале создается 00000001 а потом 1 смещается влево на i и получается 00001000)
	if b {
		return num | mask // или
	} else {
		return num &^ mask // и
	}
}

func main() {
	num := int64(18) //создаем переменную
	i := 3           //на сколько будем двигать

	num = sdvigBit(num, i, true) //устанавливаем i элемент в 1
	fmt.Printf("Результат установки на 1: %d\n", num)

	num = sdvigBit(num, i, false) //устанавливаем i элемент в 1
	fmt.Printf("Результат установки на 0: %d\n", num)
}
