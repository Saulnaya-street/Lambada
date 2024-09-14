package main

import "fmt"

//Задача 4: Встраивание структур "Работник"
//Создай структуры:
//Person с полями Name и Age,
//Employee, которая встраивает Person и добавляет поле Position (должность).
//Создай метод для структуры Employee, который выводит информацию о работнике (имя, возраст, должность).
//Создай несколько работников и выведи их информацию.

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	Position string
}

func InfoEmployee(e *Employee) {
	fmt.Printf("Имя: %s \nВозраст: %d\nДолжность: %s\n", e.Person.Name, e.Person.Age, e.Position)

}

func main() {
	employee := Employee{
		Person: Person{
			Name: "Slavik",
			Age:  20,
		},
		Position: "Boss",
	}
	InfoEmployee(&employee)
}
