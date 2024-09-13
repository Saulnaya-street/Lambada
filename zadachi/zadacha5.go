package main

import "fmt"

//Создай структуру Product, которая будет содержать следующие поля:
//Name (название товара),
//Price (цена товара),
//InStock (количество на складе).
//Напиши метод для структуры Product, который уменьшает количество товара на складе при его продаже.
//Создай несколько товаров и напиши функцию для вывода информации о каждом товаре и его наличии на складе.

type Tovary struct {
	Name  string
	Babky float64
	Kolvo int
}

func plus(z *Tovary, a int) {
	z.Kolvo = z.Kolvo - a
}

func main() {
	tovary := Tovary{
		Name:  "ганжубас",
		Babky: 50,
		Kolvo: 100,
	}
	fmt.Printf("Было товара %d. \n", tovary.Kolvo)
	plus(&tovary, 5)
	fmt.Printf("Осталось товара %d.", tovary.Kolvo)
}
