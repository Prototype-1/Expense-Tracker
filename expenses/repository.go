// repository.go - Handle data storage/retrieval
package expenses

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type ExpenseRepository struct {
	DB *sql.DB
}

func InitDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	log.Println("Connected to PostgreSQL database!")
	return db, nil
}

func NewExpenseRepository(db *sql.DB) *ExpenseRepository {
	return &ExpenseRepository{DB: db}
}

func (r *ExpenseRepository) AddExpense(description string, amount float64, category string) {
	_, err := r.DB.Exec(
		"INSERT INTO expenses (description, amount, category) VALUES ($1, $2, $3)",
		description, amount, category,
	)
	if err != nil {
		log.Printf("Error adding expense: %v", err)
	}
}

func (r *ExpenseRepository) ListExpenses() []Expense {
	rows, err := r.DB.Query("SELECT id, description, amount, date, category FROM expenses")
	if err != nil {
		log.Printf("Error retrieving expenses: %v", err)
		return nil
	}
	defer rows.Close()

	var expenses []Expense
	for rows.Next() {
		var exp Expense
		if err := rows.Scan(&exp.ID, &exp.Description, &exp.Amount, &exp.Date, &exp.Category); err != nil {
			log.Printf("Error scanning row: %v", err)
		}
		expenses = append(expenses, exp)
	}
	return expenses
}

//update expense tracker
