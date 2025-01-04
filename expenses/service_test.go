package expenses

import (
    "testing"
    "time"  
    "github.com/stretchr/testify/assert"
)

func TestAddExpense(t *testing.T) {
    mockRepo := new(MockExpenseRepository)
    
    description := "Test Expense"
    amount := 100.0
    category := "Food"
    
    mockRepo.On("AddExpense", description, amount, category).Once()
    
    mockRepo.AddExpense(description, amount, category)
    
    mockRepo.AssertExpectations(t)
}

func TestListExpenses(t *testing.T) {
    mockRepo := new(MockExpenseRepository)
    
    mockRepo.On("ListExpenses").Return([]Expense{
        {ID: 1, Description: "Test Expense 1", Amount: 50.0, Date: time.Now(), Category: "Food"},
        {ID: 2, Description: "Test Expense 2", Amount: 100.0, Date: time.Now(), Category: "Transport"},
    }).Once()
    
    expenses := mockRepo.ListExpenses()
    
    assert.Equal(t, 2, len(expenses))
    assert.Equal(t, "Test Expense 1", expenses[0].Description)
    assert.Equal(t, 50.0, expenses[0].Amount)
    mockRepo.AssertExpectations(t)
}

func TestUpdateExpenses(t *testing.T) {
    mockRepo := new(MockExpenseRepository)

    id := 1
    description := "Updated Description"
    amount := 150.0
    category := "Update Category"

    mockRepo.On("UpdateExpense", id, description, amount, category).Return(nil).Once()

    err := mockRepo.UpdateExpense(id, description, amount, category)

    assert.NoError(t, err)

    mockRepo.AssertExpectations(t)
}

func TestDeleteExpenses(t *testing.T) {
    mockRepo := new(MockExpenseRepository)

    id := 1

    mockRepo.On("DeleteExpense", id).Return(nil).Once()

    err := mockRepo.DeleteExpense(id)

    assert.NoError(t, err)

    mockRepo.AssertExpectations(t)

}