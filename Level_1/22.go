package main

//Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

import (
	"fmt"
	"math/big" //используем для корректных операций с значениями более чем int64
)

func main() {
	a := big.NewInt(1 << 20) // 2^20
	b := big.NewInt(1 << 21) // 2^21

	product := new(big.Int).Mul(a, b) // умножение
	fmt.Println("a * b =", product)

	quotient := new(big.Int).Div(a, b) // деление
	fmt.Println("a / b =", quotient)

	sum := new(big.Int).Add(a, b) // сложение
	fmt.Println("a + b =", sum)

	difference := new(big.Int).Sub(a, b) // вычитание
	fmt.Println("a - b =", difference)
}
