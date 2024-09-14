package main

import "fmt"

//Создай структуру Product, которая будет содержать следующие поля:
//Name (название товара),
//Price (цена товара),
//InStock (количество на складе).
//Напиши метод для структуры Product, который уменьшает количество товара на складе при его продаже.
//Создай несколько товаров и напиши функцию для вывода информации о каждом товаре и его наличии на складе.

type Product struct {
	Name    string
	Price   int64
	InStock int
}

func MinusInStok(p *Product, a int) {
	p.InStock = p.InStock - a
}
func InfoProduct(p *Product) {
	fmt.Printf("Название товара: %s \nцена товара: %d\nколичество: %d\n", p.Name, p.Price, p.InStock)
}

func main() {
	product := Product{
		Name:    "ганжубас",
		Price:   50,
		InStock: 100,
	}
	fmt.Printf("Было товара %d. \n", product.InStock)
	MinusInStok(&product, 5)
	fmt.Printf("Осталось товара %d. \n", product.InStock)
	InfoProduct(&product)
}
