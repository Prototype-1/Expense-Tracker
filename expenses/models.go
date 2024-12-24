// Defines data structures (e.g., Expense, Category).
package expenses

import "time"

type Expense struct {
	ID int
	Description string
	Amount float64
	Date time.Time
	Category string
}

type Category struct {
	ID int
	Name string
}