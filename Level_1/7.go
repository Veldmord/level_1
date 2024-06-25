package main

//Реализовать конкурентную запись данных в map.

import (
	"fmt"
	"sync"
)

func main() {
	mapData := make(map[string]string) //создаем мапу

	var mutex sync.Mutex //создаем мьютекс для синхронизации доступа к мапе
	numWorkers := 3      //количество горутин

	done := make(chan bool) //для синхронизации завершения горутин

	i := 1                            //играет роль ID сообщений
	for q := 0; q < numWorkers; q++ { //создаем горутины
		go func(q int) { //запускаем горутины
			for j := 0; j < 10; j++ { //делаем 10 записей в мапу
				mutex.Lock()                                                            //блокируем доступ к мапе
				mapData[fmt.Sprintf("ключ-%d-%d", j, q)] = fmt.Sprintf("Massage %d", i) //записываем данные в мапу
				i++
				mutex.Unlock() //разблокируем доступ к мапе
			}
			done <- true //сообщаем о завершении горутины
		}(q)
	}

	for q := 0; q < numWorkers; q++ { //ждем завершения всех горутин
		<-done
	}

	fmt.Println("Результат:")

	for k, v := range mapData { //проходимся по мапе и выводим результат
		fmt.Printf("key = %s  data = %s\n", k, v)
	}

	fmt.Println("Программа завершилась")

}
