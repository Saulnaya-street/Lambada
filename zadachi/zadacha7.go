package main

import "fmt"

//Задача 7: Управление автопарком
//Создай структуру Car с полями:
//Model (модель автомобиля),
//Year (год выпуска),
//Mileage (пробег).
//Напиши метод, который увеличивает пробег на заданное количество километров.
//Создай несколько автомобилей и увеличь их пробег, выводя итоговое значение.

type Tachka struct {
	Model  string
	Year   int
	Probeg int
}

func zv(z *Tachka, a int) {
	z.Probeg = z.Probeg + a
}

func main() {
	tachka := Tachka{
		Model:  "Бэха",
		Year:   2014,
		Probeg: 256,
	}
	fmt.Printf("(До) тачка %s , год %d,пробег %d. \n", tachka.Model, tachka.Year, tachka.Probeg)
	zv(&tachka, 50)
	fmt.Printf("(После) тачка %s , год %d,пробег %d. ", tachka.Model, tachka.Year, tachka.Probeg)
}
