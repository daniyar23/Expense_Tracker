/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"Expense_Tracker/internal/summary"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	description string
	amount      float64
	category    string
	date        string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "expence_tracker",
	Short: "Приложение для отслеживания расходов",
	Long: `Приложение для отслеживания расходов, позволяет добавлять, удалять и просматривать записи о расходах. Делить на
	категории, сортировать по дате, фильтровать по сумме и т.д.
	Приложение также позволяет генерировать отчеты о расходах по категориям и периодам.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Да,финансовый трекер установлен на вашем устройстве")
	},
}
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Добавить новую трату",
	RunE: func(cmd *cobra.Command, args []string) error {
		parseDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			return fmt.Errorf("неверный формат даты:%w", err)
		}
		if err := summary.AddExpense(description, amount, parseDate, category); err != nil {
			return err
		}
		fmt.Println("Трата добавлена")
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	addCmd.Flags().StringVarP(&description, "description", "d", "", "Описание траты")
	addCmd.Flags().Float64VarP(&amount, "amount", "a", 0, "Сумма траты")
	addCmd.Flags().StringVarP(&category, "category", "c", "", "Категория траты")
	addCmd.Flags().StringVarP(&date, "date", "t", "", "Дата траты (YYYY-MM-DD)")

	// делаем флаги обязательными
	addCmd.MarkFlagRequired("description")
	addCmd.MarkFlagRequired("amount")
	addCmd.MarkFlagRequired("category")
	addCmd.MarkFlagRequired("date")
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.expence_tracker.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
