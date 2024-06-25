package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func main() {
	str := "snow dog sun"               // строка
	srtSplit := strings.Split(str, " ") //разбиваем строку
	for i, j := 0, len(srtSplit)-1; i < len(srtSplit)/2; i, j = i+1, j-1 {
		srtSplit[i], srtSplit[j] = srtSplit[j], srtSplit[i] //меняем i и j эдемент
	}

	reversedStr := strings.Join(srtSplit, " ")                                         //собираем строку из масива строк
	fmt.Printf("Изначальная строка '%s' - перевернутая строка '%s'", str, reversedStr) //принтим
}
