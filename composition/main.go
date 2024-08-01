package main

import (
	"composition/store"
	"fmt"
)

func main() {
	products := map[string]store.ItemForSale{
		"Kayak": store.NewBoat("Kayak", 279, 1, false),
		"Ball":  store.NewProduct("Soccer Bal", "Soccer", 19.50),
	}
	for key, p := range products {
		switch item := p.(type) {
		case store.Describable:
			fmt.Println("Product:", item.GetName(), "Category:", item.GetCategory(), "Price:",
				item.Price(0.2))
		default:
			fmt.Println("Key:", key, "Price:", p.Price(0.2))
		}
	}

	//------------------------------------
	// rentals := []*store.RentalBoat{
	// 	store.NewRentalBoat("Rubber Ring", 10, 1, false, false, "N/A", "N/A"),
	// 	store.NewRentalBoat("Yacht", 50000, 5, true, true, "Bob", "Alice"),
	// 	store.NewRentalBoat("Super Yacht", 100000, 15, true, true, "Dora", "Charlie"),
	// }

	// for _, r := range rentals {
	// 	fmt.Printf("Rental Boat: %s, Price: %.1f, Captain: %s\n", r.Name, r.Price(0.2), r.Captain)
	// }
}
