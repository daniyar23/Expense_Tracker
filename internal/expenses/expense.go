package expenses

import (
	"errors"
	"time"
)

type Expense struct {
	ID          int
	Description string
	Amount      float64
	Date        time.Time
	Category    string
}

func NewExpense(
	id int,
	desc string,
	amount float64,
	date time.Time,
	category string,
) (*Expense, error) {

	if desc == "" {
		return nil, errors.New("Описание не может быть пустым!")
	}
	if amount <= 0 {
		return nil, errors.New("Траты должны быть больше нуля!")
	}
	if category == "" {
		return nil, errors.New("Категория не может быть пустой!")
	}
	return &Expense{
		ID:          id,
		Description: desc,
		Amount:      amount,
		Date:        date,
		Category:    category,
	}, nil
}
