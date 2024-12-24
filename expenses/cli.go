//Handles user input/output for the CLI.
package expenses

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func StartCLI() {
	repo := NewExpenseRepository()
	service := NewExpenseService(repo)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nExpense Tracker CLI")
		fmt.Println("1. Add Expense")
		fmt.Println("2. List Expenses")
		fmt.Println("3. Exit")
		fmt.Print("Please choose any from the above: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil || choice < 1 || choice > 3 {
			fmt.Println("Invalid option, please recheck!")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter description: ")
			description, _ := reader.ReadString('\n')
			description = strings.TrimSpace(description)

			fmt.Print("Enter the amount: ")
			amountInput, _ := reader.ReadString('\n')
			amountInput = strings.TrimSpace(amountInput)
			amount, err := strconv.ParseFloat(amountInput, 64)
			if err != nil || amount <= 0 {
				fmt.Println("Invalid amount. Please enter a positive number.")
				continue
			}

			fmt.Print("Enter category: ")
			category, _ := reader.ReadString('\n')
			category = strings.TrimSpace(category)

			service.AddExpense(description, amount, category)
			fmt.Println("Expense added successfully!")

		case 2:
			expenses := service.ListExpenses()
			if len(expenses) == 0 {
				fmt.Println("No expenses recorded as of now.")
			} else {
				for _, v := range expenses {
					fmt.Printf("ID: %d, Description: %s, Amount: %.2f, Date: %s, Category: %s\n",
						v.ID, v.Description, v.Amount, v.Date.Format("2006-01-02"), v.Category)
				}
			}

		case 3:
			fmt.Println("Exiting the Expense Tracker...")
			return
		}
	}
}
