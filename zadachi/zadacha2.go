package main

import "fmt"

//Задача 2: Структура "Круг"
//Создай структуру Circle, которая имеет одно поле Radius (радиус круга).
//Напиши метод Area(), который возвращает площадь круга.
//Напиши метод Perimeter(), который возвращает длину окружности.
//Формулы:
//Площадь круга: π * Radius^2
//Длина окружности: 2 * π * Radius
//Создай несколько кругов и посчитай их площадь и длину окружности.

type Circle struct {
	Radius float64
}

const Pi = 3.14

func Area(c *Circle) {
	R := Pi * c.Radius * c.Radius
	fmt.Println(R)
}
func Perimeter(c *Circle) {
	var P = 2 * Pi * c.Radius
	fmt.Println(P)
}

func main() {
	circle := Circle{
		Radius: 2.0,
	}

	Area(&circle)
	Perimeter(&circle)
	fmt.Println(circle.Radius)

}
