// Handles data storage/retrieval (in-memory for now)
package expenses

import "sync"

type ExpenseRepository struct {
	mu sync.Mutex
	expenses []Expense
}

func NewExpenseRepository() *ExpenseRepository {
	return &ExpenseRepository{}
}

func (repo *ExpenseRepository) AddExpense(expense Expense) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.expenses = append(repo.expenses, expense)
}

func (repo *ExpenseRepository) GetExpenses() []Expense {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	return repo.expenses
}