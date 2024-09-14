package main

import "fmt"

//Задача 3: Вложенные структуры "Адрес и Человек"
//Создай две структуры:
//Address с полями: City (город) и Street (улица),
//Person с полями: Name (имя), Age (возраст) и Address (вложенная структура).
//Напиши функцию, которая выводит полную информацию о человеке, включая его адрес.
//Создай несколько экземпляров структуры Person с адресами и выведи информацию о них.

type Address struct {
	Cyty   string
	Street string
}
type Person struct {
	Name string
	Age  int
	Address
}

func InfoPerson(p *Person) {
	fmt.Printf("Имя: %s \nВозраст: %d\nГород: %s\nУлица: %s\n", p.Name, p.Age, p.Address.Cyty, p.Address.Street)

}

func main() {
	person := Person{
		Name: "Sanya",
		Age:  12,
		Address: Address{
			Cyty:   "Laryak",
			Street: "Svinay",
		},
	}
	InfoPerson(&person)
}
