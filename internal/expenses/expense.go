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

var NewExpense = func(
	Id int,
	Desc string,
	Amount float64,
	Date time.Time,
	Category string,
) (*Expense, error) {

	if Desc == "" {
		return nil, errors.New("Описание не может быть пустым!")
	}
	if Amount <= 0 {
		return nil, errors.New("Траты должны быть больше нуля!")
	}
	if Category == "" {
		return nil, errors.New("Категория не может быть пустой!")
	}
	return &Expense{
		ID:          Id,
		Description: Desc,
		Amount:      Amount,
		Date:        Date,
		Category:    Category,
	}, nil
}
