package main

import "fmt"

type Payment interface {
	Pay(amount float64) string
}

type Creditcard struct {
	CardNumber string
}

type Paypal struct {
	Email string
}

func (c Creditcard) Pay(amount float64) string {
	return fmt.Sprintf("Paying $%.2f using credit card with %s  ", amount, c.CardNumber[len(c.CardNumber)-4:])
}

func (p Paypal) Pay(amount float64) string {
	return fmt.Sprintf("Paying $%.2f using email with %s ", amount, p.Email)
}
func main() {
	creditcard := Creditcard{CardNumber: "123456781234"}
	paypalemail := Paypal{Email: "rashmithreddy@gmail.com"}

	fmt.Println(creditcard.Pay(100.50))
	fmt.Println(paypalemail.Pay(100.50))
}
