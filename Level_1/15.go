package main

//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

/*
var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100] //v[:100] - это срез который ссылается на первые 100 символов строки, и содержит в себе ссылку на них, а не сами значения
}

func main() {
  someFunc()
}
*/

//Ответ: работа функции someFunc() приводит к УТЕЧКЕ ПАМЯТИ

//из за того что срез v[:100] продолжает ссылаться на огрону строку createHugeString(1 << 10),
//то сборщик мусора не может очистить выделенную память

import (
	"strings"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)       //генерируем строку размером 1024 символа
	justString = string([]byte(v[:100])) //сначала мы конверитруем строку в набор байтов []byte(v[:100]), а потом обратно в строку string(...)
	//justString = v[:100:100]
}

// генератор строки
func createHugeString(size int) string {
	return strings.Repeat("A", size)
}

func main() {
	someFunc() //вызываем функцию someFunc()
}
