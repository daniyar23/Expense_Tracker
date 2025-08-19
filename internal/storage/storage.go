package storage

import (
	"Expense_Tracker/internal/expenses"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func CreateData() error {
	datafolder := "../data"
	err := os.MkdirAll(datafolder, 0755)
	if err != nil {
		return fmt.Errorf("ошибка создания папки: %w", err)
	}

	file, err := os.Create(datafolder + "/expenses.csv")
	if err != nil {
		return fmt.Errorf("Ошибка создания файла: %w", err)
	}
	writer := csv.NewWriter(file)
	err = writer.Write([]string{"ID", "Description", "Amount", "Date", "Category"})
	if err != nil {
		return fmt.Errorf("Ошибка записи заголовков: %w", err)
	}
	writer.Flush()

	defer file.Close()
	return nil
}
func CheckData() error {
	_, err := os.Stat("data/expenses.csv")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Файл не существует")
			fmt.Println("Создаем файл...")
			err = CreateData()
			if err != nil {
				return fmt.Errorf("Ошибка создания файла: %w", err)
			}
			return nil
		}
		return fmt.Errorf("Ошибка открытия файла: %w", err)
	}

	return nil
}

func LoadData() ([]expenses.Expense, error) {
	file, err := os.Open("data/expenses.csv")
	if err != nil {
		return nil, fmt.Errorf("Ошибка открытия файла: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	if _, err := reader.Read(); err != nil {
		return nil, fmt.Errorf("Ошибка чтения заголовков: %w", err)
	}

	var tracks []expenses.Expense
	for {
		record, err := reader.Read()
		if err == io.EOF {
			fmt.Println("Файл успешно загружен")
			break
		} else if err != nil {
			return nil, fmt.Errorf("Ошибка загрузки файла: %w", err)
		}
		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, fmt.Errorf("Ошибка преобразования ID: %w", err)
		}

		amount, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return nil, fmt.Errorf("Ошибка преобразования внесенной суммы: %w", err)

		}
		date, err := time.Parse("2006-01-02", record[3])
		if err != nil {
			return nil, fmt.Errorf("Ошибка преобразования даты: %w", err)
		}
		if len(record) < 5 {
			return nil, fmt.Errorf("Недостаточно данных в строке")
		}

		tracks = append(tracks, expenses.Expense{
			ID:          id,
			Description: record[1],
			Amount:      amount,
			Date:        date,
			Category:    record[4],
		})
	}
	return tracks, nil
}

func SaveData(expenses []expenses.Expense) error {
	file, err := os.OpenFile("data/expenses.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("Ошибка создания файла: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, e := range expenses {
		row := []string{
			strconv.Itoa(e.ID),
			e.Description,
			strconv.FormatFloat(e.Amount, 'f', 2, 64),
			e.Date.Format("2006-01-02"),
			e.Category,
		}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("Ошибка записи в файл: %w", err)
		}
	}

	return nil
}
