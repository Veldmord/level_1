package main

//Реализовать собственную функцию sleep.

import (
	"fmt"
	"time"
)

func mySleep(d time.Duration) { // функция mySleep останавливает выполнение горутины на заданное время.
	<-time.After(d)
}

func main() {
	fmt.Println("Начало")

	mySleep(2 * time.Second) // Останавливаем выполнение на 2 секунды

	fmt.Println("Конец")
}
