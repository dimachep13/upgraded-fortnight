package main

import (
	"fmt"
	"net/http"

	"budget-tracker/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.ListHandler)       
	mux.HandleFunc("/add", handlers.FormHandler)   
	mux.HandleFunc("/save", handlers.SaveHandler)   

	fmt.Println("Серверзапущен!")
	fmt.Println("Откройте в браузере: http://localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Не удалось запустить сервер:", err)
	}
}