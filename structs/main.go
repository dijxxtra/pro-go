package main

import "fmt"

type Product struct {
	name, category string
	price          float64
	*Supplier
}

type Supplier struct {
	name, city string
}

// func newProduct(name, category string, price float64, supplier *Supplier) *Product {
// 	return &Product{name, category, price, supplier}
// }

// func copyProduct(product *Product) Product {
// 	p := *product
// 	s := *product.Supplier
// 	p.Supplier = &s
// 	return p
// }

func main() {
	product := Product{Supplier: &Supplier{}}

	fmt.Printf("Product name: %s, Category: %s, Price: %.2f Supplier name: %s, Supplier city: %s\n", product.name, product.category, product.price, product.Supplier.name, product.Supplier.city)

	// acme := &Supplier{"Acme Co", "New York"}
	// p1 := newProduct("Kayak", "Watersports", 275, acme)
	// p2 := copyProduct(p1)

	// p1.name = "Original Kayak"
	// p1.Supplier.name = "Boat Co"

	// for _, p := range []Product{*p1, p2} {
	// 	fmt.Println("Name:", p.name, "Supplier:", p.Supplier.name, p.Supplier.city)
	// }
}
