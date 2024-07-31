package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

type Person struct {
	name, city string
}

// ----------------------------------------------------
// type Account struct {
// 	accountNumber int
// 	expenses      []Expense
// }

// func calcTotal(expenses []Expense) (total float64) {
// 	for _, e := range expenses {
// 		total += e.getCost(true)
// 	}
// 	return
// }
//-----------------------------------------------------

func main() {
	//-----------------------------------------------------
	var expense Expense = &Product{"Kayak", "Watersport", 275}
	data := []interface{}{
		expense,
		Product{"Lifejacket", "Watersport", 48.95},
		Service{"Boat Cover", 12, 89.50, []string{}},
		Person{"Темтар Завуров", "Бабаюрт"},
		&Person{"Ислам Завуров", "Хасавюрт"},
		"This is string",
		100,
		true,
	}

	for _, item := range data {
		switch value := item.(type) {
		case Product:
			fmt.Printf("Product: %s, price: %.1f\n", value.name, value.price)
		case *Product:
			fmt.Printf("Product Pointer: %s, price: %.1f\n", value.name, value.price)
		case Service:
			fmt.Printf("Service: %s, price: %.1f\n", value.description, value.monthlyFee*float64(value.durationMonths))
		case *Service:
			fmt.Printf("Service pointer: %s, price: %.1f\n", value.description, value.monthlyFee*float64(value.durationMonths))
		case Person:
			fmt.Printf("Person: %s, city: %s\n", value.name, value.city)
		case *Person:
			fmt.Printf("Person pointer: %s, city: %s\n", value.name, value.city)
		case string, bool, int:
			fmt.Println("Built-in type:", value)
		default:
			fmt.Println("Default:", value)
		}
	}

	//-----------------------------------------------------
	// exponses := []Expense{
	// 	Service{"Boat Cover", 12, 89.50, []string{}},
	// 	Service{"Paddle Protect", 12, 8, []string{}},
	// 	&Product{"Kayak", "Watersport", 275},
	// }

	// for _, e := range exponses {
	// 	switch value := e.(type) {
	// 	case Service:
	// 		fmt.Println("Service:", value.description, "Price:", value.monthlyFee*float64(value.durationMonths))
	// 	case *Product:
	// 		fmt.Println("Product:", value.name, "Price:", value.price)
	// 	default:
	// 		fmt.Println("Exponse:", value.getName(), "Cost:", value.getCost(true))
	// 	}
	// }

	//-----------------------------------------------------
	// account := Account{
	// 	accountNumber: 12345,
	// 	expenses: []Expense{
	// 		Product{"Kayak", "Watersport", 275},
	// 		Service{"Boat Cover", 12, 89.50},
	// 	},
	// }

	// for _, e := range account.expenses {
	// 	fmt.Printf("Expense name: %s, Cost: %.2f\n", e.getName(), e.getCost(true))
	// }
	// fmt.Println("Total:", calcTotal(account.expenses))

	//-----------------------------------------------------
}
