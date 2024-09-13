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

type Redan struct {
	Radius float64
}

func Area(p *Redan, a float64) {
	R := a * p.Radius * p.Radius
	fmt.Println(R)
}
func Perimeter(p *Redan, z float64, a float64) {
	var P = z * a * p.Radius
	fmt.Println(P)
}

func main() {
	redan := Redan{
		Radius: 2.0,
	}

	Area(&redan, 3.14)
	Perimeter(&redan, 2.0, 3.14)
	fmt.Println(redan.Radius)

}
