package main

import "fmt"

//Задача 6: Управление заказами
//Создай структуру Order, которая содержит:
//Product (товар),
//Quantity (количество товара),
//TotalPrice (общая сумма заказа).
//Напиши метод для структуры Order, который считает общую сумму заказа (умножает цену товара на количество).
//Создай несколько заказов и посчитай их общую стоимость.

type Product struct {
	Name    string
	Price   int64
	InStock int
}

type Order struct {
	Product
	Quantity   int
	TotalPrice int64
}

func InfoSumOrder(o *Order) int64 {
	o.TotalPrice = (o.Product.Price * int64(o.Quantity)) / 100
	fmt.Println(o.TotalPrice)
	return o.TotalPrice
}

func main() {
	order := Order{
		Product: Product{
			Name:    "какашка",
			Price:   600,
			InStock: 60,
		},
		Quantity:   40,
		TotalPrice: 400,
	}
	InfoSumOrder(&order)
}
