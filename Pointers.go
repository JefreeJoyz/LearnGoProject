package main

import "fmt"

var msg = "Hello"

func changeMessage(msg *string) { // мы передали ссылку на адрес в памяти, где хранится это значение
	*msg += " world" // мы изменили значение в оригинальной ячейке памяти
}

func checkVariable() {
	fmt.Println(msg) // hello
	changeMessage(&msg)
	fmt.Println(msg) // hello world
}

/*
Если не использовать ссылки, то вывело бы:
hello
hello world
hello
*/

/*
var p *int    -> nil
*/
