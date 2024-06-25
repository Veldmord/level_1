package main

//Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point
//с инкапсулированными параметрами x,y и конструктором.

import (
	"fmt"
	"math"
)

type Point struct { // определение структуры Point с инкапсулированными полями x, y
	x float64
	y float64
}

func NewPoint(x, y float64) *Point { // конструктор для создания новой точки
	return &Point{x: x, y: y}
}

func (p1 *Point) Distance(p2 *Point) float64 { // метод для вычисления расстояния между двумя точками
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2)) //XY = √(xb - xa)2 + (yb - ya)2
}

func main() {
	point1 := NewPoint(1, 2) //первая точка
	point2 := NewPoint(4, 6) //вторая точка

	distance := point1.Distance(point2) // Вычисление расстояния

	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
