package main

import "fmt"

// Struct for Account Details
type AccountDetails struct {
	AccountNumber  int
	HolderName     string
	AccountBalance float64
}

// Deposit method
func (c *AccountDetails) Deposit(Amount float64) {
	c.AccountBalance += Amount
}

// Withdrawal method
func (c *AccountDetails) Withdrawal(Amount float64) {
	if Amount > c.AccountBalance {
		fmt.Println("Insufficient Balance!")
		return
	}
	c.AccountBalance -= Amount
}

// Main function
func main() {
	// Creating a map to store accounts
	M := make(map[int]*AccountDetails)

	// Adding an account
	M[1] = &AccountDetails{
		AccountNumber:  45324,
		HolderName:     "Rashmith Reddy Podduturi",
		AccountBalance: 50000,
	}

	// Display initial balance
	fmt.Println("Initial Balance:", M[1].AccountBalance)

	// Perform deposit
	M[1].Deposit(3000)
	fmt.Println("Updated Balance after Deposit:", M[1].AccountBalance)

	// Perform withdrawal
	M[1].Withdrawal(4000)
	fmt.Println("Balance after Withdrawal:", M[1].AccountBalance)

	// Attempt withdrawal beyond balance
	M[1].Withdrawal(60000) // Should print "Insufficient Balance!"
}
