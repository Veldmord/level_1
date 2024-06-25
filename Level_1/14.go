package main

//Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

import (
	"fmt"
	"reflect"
)

func detectType(str interface{}) {
	typeName := reflect.TypeOf(str) //с помощью reflect.TypeOf() определяем тип

	if typeName.Kind() == reflect.Chan { // проверяем тип на chan что бы вывести корректно результат
		fmt.Printf("Значение:Тип переменной i - канал: %s\n", typeName)
	} else {
		fmt.Printf("Значение:Тип переменной i - %v: %s\n", str, typeName)
	}
}

func main() {

	var i interface{} //создаем переменную типа interface{}

	//записываем разного типа значения в переменную и с помощью detectType() определяем тип
	i = 1
	detectType(i)

	i = "hi"
	detectType(i)

	i = true
	detectType(i)

	i = make(chan int)
	detectType(i)

	i = []int{1, 2}
	detectType(i)
}
