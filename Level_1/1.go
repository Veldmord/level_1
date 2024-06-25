package main

//Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

import "fmt"

type Human struct { // создаем структуру Human, представляющую человека
	Age  uint8  // возраст
	Name string // имя
}

type Action struct { // создаем структуру Action, представляющую действия человека
	Skill string // навык человека
	Human        // встраиваем структуру Human
}

func (h Human) printName() { // метод printName для структуры Human
	fmt.Printf("Имя нашего подопечного: %s\n", h.Name) // Вывод имени человека
}

func main() {
	human_1 := Action{ // инициализируем структуру Action
		Skill: "coding",
		Human: Human{ // инициализируем вложенную структуру Human
			Age:  19,
			Name: "Andry",
		},
	}

	human_1.printName() // вызываем метод printName через структуру Action, так как Human встроена в Action

	fmt.Printf("Его возраст: %d\n", human_1.Human.Age) // вывод возраста через вложенную структуру Human
	fmt.Printf("Его способность: %s", human_1.Skill)   // вывод навыка
}
