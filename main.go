package main

import "fmt"

func main() {

	var a, b string
	var c int
	fmt.Println("Назови свое имя...")
	fmt.Scan(&a)
	fmt.Println("Привет %i Теперь назови свою фамилию...")
	fmt.Scan(&b)
	fmt.Println("Привет,", a, b, ". Теперь назови количество полныз лет...")
	fmt.Scan(&c)
	fmt.Println("Давай подведем итог: Тебя зовут", a, ", твоя фамилия", b, "и тебе", c, "лет.")
}
