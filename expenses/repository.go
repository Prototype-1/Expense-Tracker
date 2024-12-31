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

func (r *ExpenseRepository) GetExpenseByID(id int) *Expense {
	var exp Expense

	err := r.DB.QueryRow("SELECT id, description, amount, date, category FROM expenses WHERE id = $1", id).Scan(
		&exp.ID, &exp.Description, &exp.Amount, &exp.Date, &exp.Category,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		log.Printf("Error retrieving expense with ID %d: %v", id, err)
		return nil
	}
	return &exp
}

func (r *ExpenseRepository) UpdateExpense(id int, description string, amount float64, category string) error {
	result, err := r.DB.Exec(
		"UPDATE expenses SET description = $1, amount = $2, category = $3 WHERE id = $4",
		description, amount, category, id,
	)
	if err != nil {
		log.Printf("Error updating expense with ID %d: %v", id, err)
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error fetching rows affected for update of ID %d: %v", id, err)
		return err
	}
	if rows== 0 {
		return sql.ErrNoRows 
	}
	return nil
}

func (r *ExpenseRepository) DeleteExpense(id int) error {
	result, err := r.DB.Exec("DELETE FROM expenses WHERE id = $1", id)
	if err != nil {
		log.Printf("Error deleting expense with ID %d: %v", id, err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error fetching rows affected for delete of ID %d: %v", id, err)
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

