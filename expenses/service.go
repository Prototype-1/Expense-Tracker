// service.go - Business logic
package expenses
import "fmt"


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

func (s *ExpenseService) UpdateExpense(id int, description string, amount float64, category string) error {
	existing := s.repo.GetExpenseByID(id)
	if existing == nil {
		return fmt.Errorf("expense with ID %d not found", id)
	}
	return s.repo.UpdateExpense(id, description, amount, category)
}

func (s *ExpenseService) DeleteExpense(id int) error {
	existing := s.repo.GetExpenseByID(id)
	if existing == nil {
		return fmt.Errorf("expense with ID %d not found", id)
	}
	return s.repo.DeleteExpense(id)
}


