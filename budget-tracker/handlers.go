package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"budget-tracker/internal/models" 
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Только метод GET", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("web/templates/layout.html", "web/templates/list.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки шаблонов: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "layout", models.Expenses)
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Только метод GET", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("web/templates/layout.html", "web/templates/form.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки шаблонов: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "layout", nil)
}
func SaveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Только метод POST", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Ошибка чтения формы", http.StatusBadRequest)
		return
	}

	description := r.FormValue("description")
	amountStr := r.FormValue("amount")
	date := r.FormValue("date")

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		http.Error(w, "Неверный формат суммы", http.StatusBadRequest)
		return
	}

	newExpense := models.Expense{
		Amount:      amount,
		Description: description,
		Date:        date,
	}

	models.Expenses = append(models.Expenses, newExpense)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}