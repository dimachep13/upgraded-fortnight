package models

type Expense struct {
	Amount      float64 
	Description string  
	Date        string 
}

var Expenses = []Expense{
	{Amount: 150.00, Description: "Кофе в автомате", Date: "2026-10-18"},
	{Amount: 450.50, Description: "Шаурма и сок", Date: "2026-10-18"},
}