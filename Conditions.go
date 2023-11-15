package main

import (
	"errors"
	"log"
)

func enterTheClub(age int) (string, bool) {
	if age >= 18 {
		return "come on, man \n", true
	} else {
		return "huy te, psina", false
	}
}

func enterTheClubWithErrors(age int) (string, error) {
	if age >= 18 {
		return "come on, man \n", nil
	} else {
		return "huy te, psina", errors.New("you are too young")
	}
}

func check() {
	//message, entered := enterTheClub(19)
	//fmt.Println(message, entered)

	// или так:
	//fmt.Println(enterTheClub(18))

	// или так, если нам ну нежно юзать 2й аргумент
	//message, _ := enterTheClub(10)
	//fmt.Println(message)

	// или так, если нам нужно залогировать ошибку
	message, err := enterTheClubWithErrors(10)
	if err != nil {
		log.Fatal(err)
		// fmt.Println(err) -> а так мы просто выведем текст ошибки "you are too young"
	}
	println(message)
}

// switch case logic

func Prediction(dayOfWeek string) (string, error) {
	switch dayOfWeek {
	case "пн":
		return "good monday", nil
	case "вт":
		return "good tuesday", nil
	default:
		return "default part", errors.New("invalid day of week")
	}
}
