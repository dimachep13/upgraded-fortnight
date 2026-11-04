package main

import (
	"fmt"
	"net/http"

	"budget-tracker/internal/handlers"
	"budget-tracker/internal/models"
)

func main() {

	repo, err := models.NewExpenseRepository("expenses.db")
	if err != nil {
		fmt.Println("Критическая ошибка: не удалось подключить базу данных:", err)
		return
	}
	defer repo.DB.Close() // 

	handlers.Repo = repo

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.ListHandler)
	mux.HandleFunc("/add", handlers.FormHandler)
	mux.HandleFunc("/save", handlers.SaveHandler)

	fmt.Println("=== Сервер успешно запущен с базой данных SQLite! ===")
	fmt.Println("Откройте в браузере: http://localhost:8080")

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Не удалось запустить сервер:", err)
	}
}