package main

import (
	"fmt"
	"Expense-Tracker/expenses"
	"log"
)

func main() {
	fmt.Println("Welcome to your Personal Expense Tracker")

	dataSource := "host=localhost port=5432 user=postgres password=BrainFog@1996 dbname=expense_tracker sslmode=disable"
	
	db, err := expenses.InitDB(dataSource)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	expenses.StartCLI(db)
}

