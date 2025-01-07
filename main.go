//This is main.go

package main

import (
	"fmt"
	"Expense-Tracker/expenses"
	"log"
)

func main() {
	fmt.Println("Welcome to your Personal Expense Tracker")

	dataSource := "host=localhost port=DBport user=DBusername password=DBpassword dbname=expense_tracker sslmode=disable"
	
	db, err := expenses.InitDB(dataSource)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	expenses.StartCLI(db)
}

