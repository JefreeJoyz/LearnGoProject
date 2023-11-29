package main

import (
	"fmt"
	"os"
)

func workWithEnv() {
	//path, exists := os.LookupEnv("PATH") // записываем значение PATH в переменную path
	//if exists {
	//	// Print the value of the environment variable
	//	fmt.Printf("my path is:\n %s", path)
	//}
	os.Setenv("Prod", "VarForEnv") // сохраняем значение "Prod" в переменную "VarForEnv"
	myEnv := os.Getenv("Prod")
	fmt.Printf("Gettig env from variable: %s", myEnv)
}
