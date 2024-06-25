package main

//Реализовать паттерн «адаптер» на любом примере.

import (
	"fmt"
)

type Notifier interface { // Общий интерфейс для отправки сообщений
	Send(message string) error // метод для отправки сообщений
}

type SMSNotifier struct{} // структура для отправки СМС

func (s *SMSNotifier) SendSMS(number, message string) error { // метод для отправки СМС (номер + сообщений)
	fmt.Printf("Отправка SMS на номер %s: %s\n", number, message)
	return nil
}

type EmailToSMSAdapter struct { // структура адаптера, преобразовывает email в СМС
	smsNotifier *SMSNotifier // храним объект SMSNotifier
	phoneNumber string       // храним номер телефона
}

func NewEmailToSMSAdapter(smsNotifier *SMSNotifier, phoneNumber string) *EmailToSMSAdapter { // конструктор адаптера
	return &EmailToSMSAdapter{ // возвращаем указатель на адаптер
		smsNotifier: smsNotifier,
		phoneNumber: phoneNumber,
	}
}

func (e *EmailToSMSAdapter) Send(message string) error { //реализуем метод Send для адаптера
	return e.smsNotifier.SendSMS(e.phoneNumber, message) //вызываем метод SendSMS у SMSNotifier и отправляем сообщение
}

func main() {
	smsNotifier := &SMSNotifier{}                                    //создаем объект SMSNotifier
	emailAdapter := NewEmailToSMSAdapter(smsNotifier, "+1234567890") //создаем адаптер с номером телефона

	err := emailAdapter.Send("Важное уведомление!") //отправляем email через адаптер
	if err != nil {
		fmt.Println("Ошибка при отправке уведомления:", err) //обработка ошибок
	}
}
