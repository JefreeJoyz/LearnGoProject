package main

import (
	"errors"
	"fmt"
)

var messa = [3]string{"1", "2", "3"} // массив
// var messa = []string{"1", "2", "3"} // слайс. Под капотом хранит ссылку на массив. И когда мы передаем в функцию
// слайс в качестве аргумента, то мы передаем ссылку. И дальнейшее изменение слайса в функции приведет к изменению
// слайса и за пределами функции
var messaSlice = make([]string, 5, 7) // один из способ инициализации слайса. 5 - это размер слайса, 7 - его капасити

func printMessages(messa [3]string) error {
	if len(messa) == 0 {
		return errors.New("empty array")
	}
	fmt.Println("all is fine")
	return nil
}

func workWithSlices(messaSlice []string) {
	messaSlice = append(messaSlice, "123") // добавляем в слайс элемент
	fmt.Println(len(messaSlice))
	fmt.Println(cap(messaSlice)) // показывает капасити (емкость) слайса. Если добавляем элемент в слайс, а капасити меньше,
	//	то под капотом создается копия текущего слайса в размере х2 (но не всегда именно х2), в нее копируются значения со старого, добавляется новое
	// значение, старый слайс удаляется. Операция "дорогая" по ресурсам
}
