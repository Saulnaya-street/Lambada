package main

import "fmt"

//Задача 4: Встраивание структур "Работник"
//Создай структуры:
//Person с полями Name и Age,
//Employee, которая встраивает Person и добавляет поле Position (должность).
//Создай метод для структуры Employee, который выводит информацию о работнике (имя, возраст, должность).
//Создай несколько работников и выведи их информацию.

type Borov struct {
	Name string
	Age  int
}

type Pizdec struct {
	Borov
	Position string
}

func main() {
	pizdec := Pizdec{
		Borov: Borov{
			Name: "Slavik",
			Age:  20,
		},
		Position: "Boss",
	}
	fmt.Printf("Имя %s годы жизни %d должность %s. /n ", pizdec.Name, pizdec.Age, pizdec.Position)
}
