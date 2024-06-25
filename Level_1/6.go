package main

//Реализовать все возможные способы остановки выполнения горутины.

//реализовываем отсановку горутин с помощтю контекста (отправляем сигнал остановки)

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//////с использованием контекста
	ctx, cancel := context.WithCancel(context.Background()) //используем contect для управления временем жизни и подачи сигнала отмены
	defer cancel()                                          //гарантирует отмену контекста и горутин
	go func(ctx context.Context) {                          //запускаем горутину и передаем в нее контекст
		fmt.Println("Горутина запущена")
		for {
			select {
			case <-ctx.Done(): //принимаем сигнал отановки
				fmt.Println("Горутина остановлена")
				return
			default:
				fmt.Println("Горутина работает ...") //опять работать
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(2 * time.Second)
	cancel() //отменяем контекст
	time.Sleep(1 * time.Second)
	fmt.Println("Основная функця завершена")

	////////////////////////////////////////////////////////////////////
	///////с использованием канала
	chanMess := make(chan int) //создаем канал
	i := 1

	go func() { //запускаем горутинку
		for { //бесконечный цикл
			select {
			case <-chanMess: //ловим сигнал о закрытии канала
				fmt.Println("Завершение горутины")
				return
			default: //горутина работает пока не будет сообщени в chanMess
				fmt.Printf("Работа горутины %d\n", i)
			}
			i++
		}
	}()

	time.Sleep(2 * time.Second)
	close(chanMess) //закрываем канал и следовательно останавливаем работу горутины
	time.Sleep(1 * time.Second)
	fmt.Println("Завершение работы программы")
}
