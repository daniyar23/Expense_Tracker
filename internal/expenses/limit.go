package expenses

import "errors"

type Limit struct {
	Limit float64
}

func NewLimit(limit float64) (*Limit, error) {
	if limit <= 0 {
		return nil, errors.New("Лимит должен быть больше нуля!")
	}
	return &Limit{
		Limit: limit,
	}, nil
}
