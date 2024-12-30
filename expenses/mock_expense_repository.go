package expenses

import "github.com/stretchr/testify/mock"

type MockExpenseRepository struct {
    mock.Mock
}

func (m *MockExpenseRepository) AddExpense(description string, amount float64, category string) {
    m.Called(description, amount, category)
}

func (m *MockExpenseRepository) ListExpenses() []Expense {
    args := m.Called()
    return args.Get(0).([]Expense)
}

