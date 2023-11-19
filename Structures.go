package main

import "fmt"

type User struct {
	name string
	age  int
	sex  string
}

// создаем метод с "типом структуры", к которому можно обращаться с любой инициализированной ранее структуры.
// в нашем случае, мы в функции testStruct инициализировали user. И уже с него мы вызвали метод getName()
// в свифте эти методы были внутри структуры, а здесь снаружи и с некоторой пометкой
func (u User) getName() {
	println("func (u User)")
}

func testStruct() {
	user := User{
		name: "Jack",
		age:  34,
		sex:  "Male",
	}

	NewDB()

	user.getName()

	fmt.Println(user.sex)
	fmt.Printf("%+v", user) // выводит вместе с названиями полей
}

type DB struct {
	m map[string]string
}

// NewDB - Конструктор, который инициализирует нашу структуру и теперь мы можем писать в нее что то
func NewDB() *DB {
	return &DB{
		m: make(map[string]string),
	}
}
