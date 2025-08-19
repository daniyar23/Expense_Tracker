package summary

import (
	"Expense_Tracker/internal/expenses"
	"Expense_Tracker/internal/storage"
	"fmt"
	"time"
)

func AddExpense(description string,
	amount float64,
	date time.Time,
	category string) error {
	expense, err := storage.LoadData()
	if err != nil {
		return fmt.Errorf("Error:%w", err)
	}
	newexpense, err := expenses.NewExpense(len(expense)+1, description, amount, date, category)
	if err != nil {
		return err
	}
	expense = append(expense, *newexpense)
	err = storage.SaveData(expense)
	if err != nil {
		return err
	}
	return nil
}
