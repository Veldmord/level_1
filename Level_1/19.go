package main

import "fmt"

//Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

func ReversStr(str string) string {
	runeStr := []rune(str)                                               //преобразовываем строку в массив рун
	for i, j := 0, len(runeStr)-1; i < len(runeStr)/2; i, j = i+1, j-1 { //
		runeStr[i], runeStr[j] = runeStr[j], runeStr[i] // меняем i и j символы местами
	}
	return string(runeStr) //возвращаем перевернутую строку
}

func main() {
	var str string
	fmt.Println("Введите строку:")
	fmt.Scan(&str) //захватываем из консоли строку

	revStr := ReversStr(str) //переворачиваем

	fmt.Print(revStr) //принтим
}
