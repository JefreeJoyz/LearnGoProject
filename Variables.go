package main

import "fmt"

// глобальные переменные
var var1 = "yo"
var var2 = 5

// Константа
const var3 bool = true

func main2() {
	fmt.Println("hello world")
	test1()
}

func test1() {
	// Локальные переменные
	a := "aaa"
	b := 42
	c := false
	fmt.Printf("Переменная 1: %v, переменная 2: %v \n", var1, var2)
	fmt.Println(a, b, c)
}
