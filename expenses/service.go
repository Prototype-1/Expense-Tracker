// Contains core business logic
package expenses

import "time"

type ExpenseService struct {
	repo *ExpenseRepository
}

func NewExpenseService(repo *ExpenseRepository) *ExpenseService {
	return &ExpenseService{repo: repo}
}

func (s *ExpenseService) AddExpense(description string, amount float64, category string) {
	expense := Expense {
		ID: len(s.repo.GetExpenses()),
		Description: description,
		Amount: amount,
		Date: time.Now(),
		Category: category,
	}
	s.repo.AddExpense(expense)
}

func (s *ExpenseService) ListExpenses() []Expense {
	return s.repo.GetExpenses()
}