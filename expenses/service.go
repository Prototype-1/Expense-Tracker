// service.go - Business logic
package expenses


type ExpenseService struct {
	repo *ExpenseRepository
}

func NewExpenseService(repo *ExpenseRepository) *ExpenseService {
	return &ExpenseService{repo: repo}
}

func (s *ExpenseService) AddExpense(description string, amount float64, category string) {
	expense := Expense{
		Description: description,
		Amount:      amount,
		Category:    category,
	}
	s.repo.AddExpense(expense.Description, expense.Amount, expense.Category)
}

func (s *ExpenseService) ListExpenses() []Expense {
	return s.repo.ListExpenses()
}
//Need to add edit and delete functions
