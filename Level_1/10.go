package main

//10.	Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
//Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

//Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

import (
	"fmt"
)

func groupTemp(tempArr []float64) map[int][]float64 {
	mapTemp := make(map[int][]float64) //формируем мапу со слайсом
	for _, temp := range tempArr {     //проходимся по массиву
		value := int(temp/10) * 10                    //формируем шаг/ключ
		mapTemp[value] = append(mapTemp[value], temp) //добовляем значени из массива в мапу
	}
	return mapTemp //возвращаем мапу
}

func main() {
	tempArr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5} //наш массив

	groupTemps := groupTemp(tempArr) //отправляем группироваться

	fmt.Println(groupTemps) //смотрим
}
