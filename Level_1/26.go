package main

//Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

//Например:
//abcd — true
//abCdefAaf — false
//aabcd — false

import (
	"fmt"
	"strings"
)

// Функция проверяет, все ли символы в строке уникальны (без учета регистра).
func isUnique(str string) bool {
	// Преобразуем строку в нижний регистр для регистронезависимой проверки
	str = strings.ToLower(str)

	// Используем map для отслеживания встречавшихся символов
	seenChars := make(map[rune]bool)

	// Проходим по всем символам строки
	for _, char := range str {
		// Если символ уже встречался, возвращаем false
		if seenChars[char] {
			return false
		}
		// Иначе помечаем символ как встречавшийся
		seenChars[char] = true
	}

	// Все символы уникальны
	return true
}

func main() {
	fmt.Println(isUnique("abcd"))      // true
	fmt.Println(isUnique("abCdefAaf")) // false
	fmt.Println(isUnique("aabcd"))     // false
	fmt.Println(isUnique("123456"))    // true
	fmt.Println(isUnique("123456&*"))  // true
	fmt.Println(isUnique("123456&*&")) // false
}
