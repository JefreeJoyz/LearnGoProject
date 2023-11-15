package main

import "fmt"

// Функция, которая принимает стрингу и возвращает тоже стрингу
func mainFunc(arg1 string) string {
	//return "Привет, " + arg1
	fmt.Println("Привет, " + arg1)
	return arg1
}

func toStringFunc(name string, age int) string {
	return fmt.Sprintf("Привет, %s! Тебе %d годика", name, age) // Sprintf конвертирует все в строку
}

func showMessage() {
	myMessage := toStringFunc("Jack", 34)
	fmt.Println(myMessage)
}

func findMin(numbers ...int) int { // эта запись означает, что в эту функцию можно передать
	// неограниченное кол-во значений int

	// do some logic
	min := 0
	return min
}
