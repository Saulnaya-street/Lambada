package main

import "fmt"

//Задача 6: Управление заказами
//Создай структуру Order, которая содержит:
//Product (товар),
//Quantity (количество товара),
//TotalPrice (общая сумма заказа).
//Напиши метод для структуры Order, который считает общую сумму заказа (умножает цену товара на количество).
//Создай несколько заказов и посчитай их общую стоимость.

type Order struct {
	Product string
	Kolvo   int
	Obchag  float64
}

func minus(z *Order, a float64) float64 {
	c := a * float64(z.Kolvo)
	fmt.Println(c)
	return c
}

func main() {
	order := Order{
		Product: "ганжубас",
		Kolvo:   50,
		Obchag:  100.0,
	}
	minus(&order, 50.0)
}
