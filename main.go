package main

import (
	"Expense_Tracker/cmd"
	"Expense_Tracker/internal/summary"
	"time"
)

func main() {
	summary.AddExpense("ice eqeqcream", 123, time.Now(), "Food")
	cmd.Execute()
}
