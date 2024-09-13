package main

import "fmt"

//Задача 1: Простая структура "Книга"
//Создай структуру Book, которая содержит следующие поля:
//Title (название книги),
//Author (автор),
//Pages (количество страниц).
//Напиши функцию, которая принимает структуру Book и выводит информацию о книге в виде строки.
//Создай несколько экземпляров структуры Book и выведи их на экран.

type Book struct {
	Title  string
	Author string
	Pages  int
}

func main() {
	niga := Book{
		Title:  "пальчик чешится ай ай ай",
		Author: "Кузьмин А.Е",
		Pages:  2,
	}

	fmt.Println(niga.Title)
	fmt.Println(niga.Author)
	fmt.Println(niga.Pages)

}
