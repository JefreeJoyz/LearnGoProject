package main

import (
	"fmt"
	"learn-go-project/services"
	"net/http"
)

func main() {

	//mainFunc("Мир")
	//showMessage()
	//enterTheClub(19)
	//check()
	//checkVariable()
	//printMessages(messa)
	//workWithSlices(messaSlice)
	//fForMatrix()
	//PanicRecoverDefer()
	//testMap()
	//testStruct()
	//testInterface()
	//services.GetMyIp()
	//testUnitTest()
	http.HandleFunc("/handler1", services.GetMyIpEndpoint)
	http.ListenAndServe(":8080", nil)
	fmt.Println("hello")
}
