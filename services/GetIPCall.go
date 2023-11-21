package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetMyIp() {
	url := "https://api.ipify.org?format=json"
	var getIP GetIP
	//var getDetailsIP GetDetailsIP

	body := GetData(url)

	// Разбираем JSON и заполняем структуру
	err := json.Unmarshal(body, &getIP)
	if err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return
	}

	// Используем данные
	fmt.Printf("My ip is %s\n", getIP.Ip)
	myIP := getIP.Ip
	GetDetailIP(myIP)

}

func GetDetailIP(ip string) {
	url := fmt.Sprintf("https://ipinfo.io/%s/geo", ip)
	var getDetailsIP GetDetailsIP

	body := GetData(url)

	// Разбираем JSON и заполняем структуру
	err := json.Unmarshal(body, &getDetailsIP)
	if err != nil {
		fmt.Println("Ошибка при разборе JSON in GetDetailIP:", err)
		return
	}

	// Используем данные
	fmt.Println(getDetailsIP.Country)

}

func GetData(url string) []byte {
	// Отправляем GET-запрос
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("There is some errors with trying to get url")
		return nil
	}
	defer response.Body.Close()

	// Чтение ответа
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("There is some errors with trying to read response")
		return nil
	}
	return body
}
