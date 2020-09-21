package types

// Money represents a monetary amount in minimum units (cents, kopecks, diramas, etc.).
type Money int64

// Category передставляет собой категорию, в которой был совершён платёж (авто, аптеки, рестораны и т.д.).
type Category string

// Payment presents information about the payment.
type Payment struct {
	ID 			int
	Amount 		Money
	Category 	Category
}