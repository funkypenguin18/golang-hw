package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse" // Импортируем пакет для переворота строки
)

func main() {
	// Исходная строка
	originalString := "Hello, OTUS!"

	// Переворачиваем строку
	reversed := reverse.String(originalString) // reverse.Reverse принимает строку, преобразует её в руны
	// и переворачивает(возвращает свой аргумент string в обратном руническом порядке слева направо) её.
	// Возвращает строку

	// Печатаем перевернутую строку
	fmt.Println(reversed)
}
