package main

import "fmt"

func testMap() {
	users := map[string]int{
		"Vova":   15,
		"Kostya": 30,
		"Anton":  40,
	}
	age, exists := users["Vov2a"] // exists - булевая переменная, отвечающая за то, существует ли элемент по такому ключу
	if exists {
		fmt.Println("Vova существует", age)
	} else {
		fmt.Println("чет пошло не так")
	}

	// Добавление значение в мапу
	users["Serega"] = 21

	// Удаление значения из мапы
	delete(users, "Vova")

	// Можем итерировать мапу
	for key, value := range users {
		fmt.Println(key, value)
	}

	//fmt.Println(users["Vova"])
}
