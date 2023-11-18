package main

import "fmt"

func PanicRecoverDefer() {
	//panic("heeeelp") -> падает приложение. Функция recover() позволяет не уронить (panic) приложение
	testDefer()
}

func testDefer() {
	defer pr("message 1") // ключевое слово defer означает, что мы откладываем выполнение этой функции
	// до тех пор, пока не будет выполнено все остальное
	pr("message 2")
	pr("message 3")
	// функция выведет 2, 3, 1
}

func pr(message string) {
	fmt.Println(message)
}
