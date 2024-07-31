package main

import (
	"fmt"
	. "packages/fmt"
	"packages/store"
	"packages/store/cart"
)

func main() {

	product := store.NewProduct("Kayak", "Watersport", 279)
	cart := cart.Cart{
		CustomerName: "Alice",
		Products:     []store.Product{*product},
	}

	fmt.Println("Name:", cart.CustomerName)
	fmt.Println("Total:", ToCurrency(cart.GetTotal()))
}
