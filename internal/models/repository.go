package models

import (
	"database/sql"
	"fmt" 
	"strings"
	_ "modernc.org/sqlite"
)

type ExpenseRepository struct {
	DB *sql.DB
}

func NewExpenseRepository(dbPath string) (*ExpenseRepository, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE
	);`)
	if err != nil {
		return nil, err
	}

	var count int
	db.QueryRow("SELECT COUNT(*) FROM categories").Scan(&count)
	if count == 0 {
		categories := []string{"Еда", "Транспорт", "Жилье", "Развлечения", "Прочее"}
		for _, cat := range categories {
			db.Exec("INSERT INTO categories (name) VALUES (?)", cat)
		}
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS expenses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		amount REAL NOT NULL,
		description TEXT NOT NULL,
		date TEXT NOT NULL,
		category_id INTEGER,
		FOREIGN KEY (category_id) REFERENCES categories(id)
	);`)
	if err != nil {
		return nil, err
	}

	return &ExpenseRepository{DB: db}, nil
}
func (r *ExpenseRepository) GetCategories() ([]Category, error) {
	rows, err := r.DB.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Category
	for rows.Next() {
		var c Category
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return nil, err
		}
		list = append(list, c)
	}
	return list, nil
}

func (r *ExpenseRepository) Create(e Expense) error {
	_, err := r.DB.Exec(
		"INSERT INTO expenses (amount, description, date, category_id) VALUES (?, ?, ?, ?)",
		e.Amount, e.Description, e.Date, e.CategoryID,
	)
	return err
}

func (r *ExpenseRepository) GetFiltered(categoryID string, fromDate string, toDate string) ([]Expense, error) {

	query := `
		SELECT e.id, e.amount, e.description, e.date, e.category_id, IFNULL(c.name, 'Без категории') 
		FROM expenses e
		LEFT JOIN categories c ON e.category_id = c.id
		WHERE 1=1
	`
	var args []interface{}

	if categoryID != "" {
		query += " AND e.category_id = ?"
		args = append(args, categoryID)
	}
	if fromDate != "" {
		query += " AND e.date >= ?"
		args = append(args, fromDate)
	}
	if toDate != "" {
		query += " AND e.date <= ?"
		args = append(args, toDate)
	}

	query += " ORDER BY e.id DESC"

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Expense
	for rows.Next() {
		var e Expense
		err := rows.Scan(&e.ID, &e.Amount, &e.Description, &e.Date, &e.CategoryID, &e.CategoryName)
		if err != nil {
			return nil, err
		}
		list = append(list, e)
	}
	return list, nil
}
func (r *ExpenseRepository) GetByID(id int) (Expense, error) {
	var e Expense
	err := r.DB.QueryRow(
		"SELECT id, amount, description, date, category_id FROM expenses WHERE id = ?",
		id,
	).Scan(&e.ID, &e.Amount, &e.Description, &e.Date, &e.CategoryID)
	return e, err
}

func (r *ExpenseRepository) Update(e Expense) error {
	_, err := r.DB.Exec(
		"UPDATE expenses SET amount = ?, description = ?, date = ?, category_id = ? WHERE id = ?",
		e.Amount, e.Description, e.Date, e.CategoryID, e.ID,
	)
	return err
}

func (r *ExpenseRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM expenses WHERE id = ?", id)
	return err
}