package main

// //1.
// func printPrice() {
// 	kayakPrice := 275.00
// 	kayakTax := kayakPrice * 0.2
// 	fmt.Println("Price:", kayakPrice, "Tax:", kayakTax)
// }

// //2.
// func printPrice(product string, price float64, taxRate float64) {
// 	taxAmount := price * taxRate
// 	fmt.Println(product, "price:", price, "Tax:", taxAmount)
// }

// //3.
// func printSuppliers(product string, suppliers []string) {
// 	for _, supplier := range suppliers {
// 		fmt.Println("Product:", product, "Supplier:", supplier)
// 	}
// }

// //3.1
// func printSuppliers(product string, suppliers ...string) {
// 	for _, supplier := range suppliers {
// 		fmt.Println("Product:", product, "Supplier:", supplier)
// 	}
// }

// //3.2, 3.3
// func printSuppliers(product string, suppliers ...string) {
// 	if len(suppliers) == 0 {
// 		fmt.Println("Product:", product, "Suppliers: (none)")
// 	} else {
// 		for _, supplier := range suppliers {
// 			fmt.Println("Product:", product, "Suppliers:", supplier)
// 		}
// 	}
// }

// // 4.
// func swapValues(first, second int) {
// 	fmt.Println("-- Before swap:", first, second)
// 	first, second = second, first
// 	fmt.Println("-- After swap:", first, second)
// }

// // 4.1
// func swapValues(first, second *int) {
// 	fmt.Println("-- Before swap:", *first, *second)
// 	*first, *second = *second, *first
// 	fmt.Println("-- After swap:", *first, *second)
// }

// //5.
// func calcTax(price float64) float64 {
// 	return price + (price * 0.2)
// }

// //6.
// func swapValues(first, second int) (int, int) {
// 	return second, first
// }

// //7.
// func calcTax(price float64) float64 {
// 	if price > 100 {
// 		return price * 0.2
// 	}
// 	return -1
// }

// //7.1
// func calcTax(price float64) (float64, bool) {
// 	if price > 100 {
// 		return price * 0.2, true
// 	}
// 	return 0, false
// }

// //8.
// func calcTax(price float64) (float64, bool) {
// 	if price > 100 {
// 		return price * 0.2, true
// 	}
// 	return 0, false
// }

// func calcTotalPrice(products map[string]float64, minSpend float64) (total, tax float64) {
// 	total = minSpend
// 	for _, price := range products {
// 		if taxAmount, due := calcTax(price); due {
// 			total += taxAmount
// 			tax += taxAmount
// 		} else {
// 			total += price
// 		}
// 	}
// 	return
// }

// //9.
// func calcTotalPrice(products map[string]float64) (count int, total float64) {
// 	count = len(products)
// 	for _, price := range products {
// 		total += price
// 	}
// 	return
// }

// //10.
// func calcTotalPrice(products map[string]float64) (count int, total float64) {
// 	fmt.Println("Function started")
// 	defer fmt.Println("First defer call")
// 	count = len(products)
// 	for _, price := range products {
// 		total += price
// 	}
// 	defer fmt.Println("Second defer call")
// 	fmt.Println("Function about to return")
// 	return
// }

func main() {
	// //1.Определение простой функции
	// fmt.Println("About to call function")
	// printPrice()
	// fmt.Println("Function complete")

	// //2. Определение и использование параметров функции
	// printPrice("Kayak", 275.00, 0.2)
	// printPrice("Lifejacket", 48.95, 0.2)
	// printPrice("Soccer Ball", 19.5, 0.15)

	// //3. Определение вариационных параметров
	// printSuppliers("Kayak", []string{"Acme Kayaks", "Bob's Boats", "Crazy Canoes"})
	// printSuppliers("Lifejacket", []string{"Sail Safe Co"})

	// //3.1
	// printSuppliers("Kayak", "Acme Kayaks", "Bob's Boat", "Crazy Canoes")
	// printSuppliers("Lifejacket", "Sail Safe Co")

	// //3.2
	// printSuppliers("Kayak", "Acme Kayaks", "Bob's Boats", "Crazy Canoes")
	// printSuppliers("Lifejacket", "Sail Safe Co")
	// printSuppliers("Soccer Ball")

	// //3.3
	// names := []string{"Acme Kayaks", "Bob's Boats", "Crazy Canoes"}
	// printSuppliers("Kayak", names...)
	// printSuppliers("Lifejacket", "Sail Safe Co")
	// printSuppliers("Soccer Ball")

	// //4.
	// val1, val2 := 10, 20
	// fmt.Println("Before calling function:", val1, val2)
	// swapValues(val1, val2)
	// fmt.Println("After calling function:", val1, val2)

	// //4.1
	// val1, val2 := 10, 20
	// fmt.Println("Before calling function:", val1, val2)
	// swapValues(&val1, &val2) // Передаем указатели
	// fmt.Println("After calling function:", val1, val2)

	// //5. Определение и использование результатов функции
	// products := map[string]float64{
	// 	"Kayak":      275,
	// 	"Lifejacket": 48.95,
	// }

	// for product, price := range products {
	// 	fmt.Println("Product:", product, "Price with Tax:", calcTax(price))
	// }

	// //6. Возврат функцией нескольких результатов
	// val1, val2 := 10, 20
	// fmt.Println("Before calling function:", val1, val2)
	// val1, val2 = swapValues(val1, val2)
	// fmt.Println("After calling function:", val1, val2)

	// //7. Использование нескольких результатов вместо нескольких значений
	// products := map[string]float64{
	// 	"Kayak":      275,
	// 	"Lifejacket": 48.95,
	// }
	// for product, price := range products {
	// 	tax := calcTax(price)
	// 	if tax != -1 {
	// 		fmt.Println("Product:", product, "Tax:", tax)
	// 	} else {
	// 		fmt.Println("Product:", product, "No tax due")
	// 	}
	// }

	// //7.1
	// products := map[string]float64{
	// 	"Kayak":      275,
	// 	"Lifejacket": 48.95,
	// }
	// for product, price := range products {
	// 	if taxAmount, taxDue := calcTax(price); taxDue {
	// 		fmt.Println("Product:", product, "Tax:", taxAmount)
	// 	} else {
	// 		fmt.Println("Product:", product, "No tax due")
	// 	}
	// }

	// //8.Использование именованных результатов
	// products := map[string]float64{
	// 	"Kayak":      275,
	// 	"Lifejacket": 48.95,
	// }
	// total1, tax1 := calcTotalPrice(products, 10)
	// fmt.Println("Total1:", total1, "Tax:", tax1)
	// total2, tax2 := calcTotalPrice(nil, 10)
	// fmt.Println("Total2:", total2, "Tax:", tax2)

	// //9. Использование пустого идентификатора для сброса результатов
	// products := map[string]float64{
	// 	"Kayak":      275,
	// 	"Lifejacket": 48.95,
	// }
	// _, total := calcTotalPrice(products)
	// fmt.Println("Total:", total)

	//10. Использование ключевого слова defer
	// products := map[string]float64{
	// 	"Kayak":      275,
	// 	"Lifejacket": 48.95,
	// }
	// _, total := calcTotalPrice(products)
	// fmt.Println("Total:", total)
}
