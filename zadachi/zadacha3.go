package main

import "fmt"

//Задача 3: Вложенные структуры "Адрес и Человек"
//Создай две структуры:
//Address с полями: City (город) и Street (улица),
//Person с полями: Name (имя), Age (возраст) и Address (вложенная структура).
//Напиши функцию, которая выводит полную информацию о человеке, включая его адрес.
//Создай несколько экземпляров структуры Person с адресами и выведи информацию о них.

type OPG struct {
	Cyty   string
	Street string
}
type Kaban struct {
	Name string
	Age  int
}

func DayZ(z *OPG, x *OPG, c *Kaban, v *Kaban) {
	fmt.Println(z.Cyty, x.Street, c.Name, v.Age)

}

func main() {
	opg := OPG{
		Cyty:   "Laryak",
		Street: "Svinay",
	}
	kaban := Kaban{
		Name: "Danil",
		Age:  27,
	}
	DayZ(&opg, &opg, &kaban, &kaban)
}
