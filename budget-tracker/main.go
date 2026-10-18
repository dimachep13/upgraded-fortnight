
package main

import (
	"fmt"
	"net/http"
)
func indexHandler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprint(writer, "Привет! Это главная страница учета расходов.")
}
func infoHandler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprint(writer, "Семестровая работа по Go. Тут будет учет личных трат.")
}
func pingHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(writer, "Ошибка 405: разрешен только GET запрос")
		return
	}
	fmt.Fprint(writer, "pong")
}
func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/about", infoHandler)
	router.HandleFunc("/ping", pingHandler)

	fmt.Println("Сервер работает на порту 8080...")
	
	http.ListenAndServe(":8080", router)
}