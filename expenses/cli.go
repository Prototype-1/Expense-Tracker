// cli.go - Handles user input/output for the CLI
package expenses

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"database/sql"
)

func StartCLI(db *sql.DB) {
	repo := NewExpenseRepository(db) 
	service := NewExpenseService(repo)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nExpense Tracker CLI")
		fmt.Println("1. Add Expense")
		fmt.Println("2. List Expenses")
		fmt.Println("3. Update Expense")
		fmt.Println("4. Delete Expense")
		fmt.Println("5. Exit")
		fmt.Print("Please choose an option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil || choice < 1 || choice > 5 {
			fmt.Println("Invalid option, please try again!")
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
			fmt.Print("Enter the ID of the expense to update: ")
			idInput, _ := reader.ReadString('\n')
			idInput = strings.TrimSpace(idInput)
			id, err := strconv.Atoi(idInput)
			if err != nil || id <= 0 {
				fmt.Println("Invalid ID. Please enter a valid numeric ID.")
				continue
			}

			fmt.Print("Enter new description: ")
			description, _ := reader.ReadString('\n')
			description = strings.TrimSpace(description)

			fmt.Print("Enter new amount: ")
			amountInput, _ := reader.ReadString('\n')
			amountInput = strings.TrimSpace(amountInput)
			amount, err := strconv.ParseFloat(amountInput, 64)
			if err != nil || amount <= 0 {
				fmt.Println("Invalid amount. Please enter a positive number.")
				continue
			}

			fmt.Print("Enter new category: ")
			category, _ := reader.ReadString('\n')
			category = strings.TrimSpace(category)

			err = service.UpdateExpense(id, description, amount, category)
			if err != nil {
				fmt.Printf("Error updating expense: %v\n", err)
			} else {
				fmt.Println("Expense updated successfully!")
			}

		case 4: 
			fmt.Print("Enter the ID of the expense to delete: ")
			idInput, _ := reader.ReadString('\n')
			idInput = strings.TrimSpace(idInput)
			id, err := strconv.Atoi(idInput)
			if err != nil || id <= 0 {
				fmt.Println("Invalid ID. Please enter a valid numeric ID.")
				continue
			}

			err = service.DeleteExpense(id)
			if err != nil {
				fmt.Printf("Error deleting expense: %v\n", err)
			} else {
				fmt.Println("Expense deleted successfully!")
			}

		case 5: 
			fmt.Println("Exiting the Expense Tracker...")
			return
		}
	}
}




