//A CLI application, or command line interface application.
package main
import (
"fmt"
"Expense-Tracker/expenses"
)

func main() {
	fmt.Println("Welcome to the Personal Expense Tracker")
	expenses.StartCLI()
}
