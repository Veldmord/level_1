package main

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10} // подопытные

	var wg sync.WaitGroup // создаем ВГ для ожидания выполнения всех горутин
	sum := 0              // сюда будем суммировать

	for _, i := range arr { //проходимся по циклу
		wg.Add(1)          //увеличиваем счетчик горутин на 1
		go func(num int) { //запускаем горутину и передаем значение из массива
			defer wg.Done()  //гарантированно уменьшаем счетчик на 1
			sum += num * num //прибавляем еще 1 квадрат в наш общаг
		}(i) //вызов анонимной функции
	}

	wg.Wait() // ожидание завершение всех горутин

	fmt.Println(sum) //вывод: 220
}
