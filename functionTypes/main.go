package main

//1.
// func calcWithTax(price float64) float64 {
// 	return price + (price * 0.2)
// }

// func calcWithoutTax(price float64) float64 {
// 	return price
// }

// //2.
// func calcWithTax(price float64) float64 {
// 	return price + (price * 0.2)
// }

// func calcWithoutTax(price float64) float64 {
// 	return price
// }

// func printPrice(product string, price float64, calculator func(float64) float64) {
// 	fmt.Println("Product:", product, "Price:", calculator(price))
// }

// //3.
// func calcWithTax(price float64) float64 {
// 	return price + (price * 0.2)
// }

// func calcWithoutTax(price float64) float64 {
// 	return price
// }

// func printPrice(product string, price float64, calculator func(float64) float64) {
// 	fmt.Println("Product:", product, "Price:", calculator(price))
// }

// func selectCalculator(price float64) func(float64) float64 {
// 	if price > 100 {
// 		return calcWithTax
// 	}
// 	return calcWithoutTax
// }

// //4.
// type calcFunc func(float64) float64

// func calcWithTax(price float64) float64 {
// 	return price + (price * 0.2)
// }

// func calcWithoutTax(price float64) float64 {
// 	return price
// }

// func printPrice(product string, price float64, calculator calcFunc) {
// 	fmt.Println("Product:", product, "Price:", calculator(price))
// }

// func selectCalculator(price float64) calcFunc {
// 	if price > 100 {
// 		return calcWithTax
// 	}
// 	return calcWithoutTax
// }

// // //5.
// type calcFunc func(float64) float64

// func printPrice(product string, price float64, calculator calcFunc) {
// 	fmt.Println("Product:", product, "Price:", calculator(price))
// }

// func selectCalculator(price float64) calcFunc {
// 	if price > 100 {
// 		// return func(price float64) float64 {
// 		// 	return price + (price * 0.2)
// 		// }
// 		var withTax calcFunc = func(price float64) float64 {
// 			return price + (price * 0.2)
// 		}
// 		// withTax := func(price float64) float64 {
// 		// 	return price + (price * 0.2)
// 		// }
// 		return withTax
// 	} else {
// 		// return func(price float64) float64 {
// 		// 	return price
// 		// }
// 		var withoutTax calcFunc = func(price float64) float64 {
// 			return price
// 		}
// 		return withoutTax
// 	}
// }

// //6.
// type calcFunc func(float64) float64

// func printPrice(product string, price float64, calculator calcFunc) {
// 	fmt.Println("Product:", product, "Price:", calculator(price))
// }

// //6.1
// type calcFunc func(float64) float64

// func printPrice(product string, price float64, calculator calcFunc) {
// 	fmt.Println("Product:", product, "Price:", calculator(price))
// }

// func priceCalcFactory(threshold, rate float64) calcFunc {
// 	return func(price float64) float64 {
// 		if price > threshold {
// 			return price + (price * rate)
// 		}
// 		return price
// 	}
// }

// //7.
// type calcFunc func(float64) float64

// func printPrice(product string, price float64, calculator calcFunc) {
// 	fmt.Println("Product:", product, "Price:", calculator(price))
// }

// var prizeGiveaway = false

// func priceCalcFactory(threshold, rate float64) calcFunc {
// 	// 7.1 fixedPrizeGiveaway := prizeGiveaway
// 	return func(price float64) float64 {
// 		// 7.1 if fixedPrizeGiveaway {
// 		if prizeGiveaway {
// 			return 0
// 		} else if price > threshold {
// 			return price + (price * rate)
// 		}
// 		return price
// 	}
// }

func main() {
	// //1. Понимание типов функций
	// products := map[string]float64{
	// 	"Kayak":      275,
	// 	"Lifejacket": 48.95,
	// }
	// for product, price := range products {
	// 	var calcFunc func(float64) float64
	// 	if price > 100 {
	// 		calcFunc = calcWithTax
	// 	} else {
	// 		calcFunc = calcWithoutTax
	// 	}
	// 	totalPrice := calcFunc(price)
	// 	fmt.Println("Product:", product, "Price:", totalPrice)
	// }

	// //2. Использование функций в качестве аргументов
	// products := map[string]float64{
	// 	"Kayak":      275,
	// 	"Lifejacket": 48.95,
	// }

	// for product, price := range products {
	// 	if price > 100 {
	// 		printPrice(product, price, calcWithTax)
	// 	} else {
	// 		printPrice(product, price, calcWithoutTax)
	// 	}
	// }

	// //3. Использование функций в качестве результатов
	// products := map[string]float64{
	// 	"Kayak":     275,
	// 	"Lifejaket": 48.95,
	// }
	// for product, price := range products {
	// 	printPrice(product, price, selectCalculator(price))
	// }

	// //4. Создание псевдонимов функциональных типов
	// products := map[string]float64{
	// 	"Kayak":     275,
	// 	"Lifejaket": 48.95,
	// }
	// for product, price := range products {
	// 	printPrice(product, price, selectCalculator(price))
	// }

	//5. Использование литерального синтаксиса функции
	// products := map[string]float64{
	// 	"Kayak":     275,
	// 	"Lifejaket": 48.95,
	// }
	// for product, price := range products {
	// 	printPrice(product, price, selectCalculator(price))
	// }

	// //6. Понимание замыкания функции
	// watersportsProducts := map[string]float64{
	// 	"Kayak":      275,
	// 	"Lifejacket": 48.95,
	// }
	// soccerProducts := map[string]float64{
	// 	"Soccer Ball": 19.50,
	// 	"Stadium":     79500,
	// }
	// // Создаем calc для watersports
	// calc := func(price float64) float64 {
	// 	if price > 100 {
	// 		return price + (price * 0.2)
	// 	}
	// 	return price
	// }

	// for product, price := range watersportsProducts {
	// 	printPrice(product, price, calc)
	// }
	// // Меняем calc под soccerProducts
	// calc = func(price float64) float64 {
	// 	if price > 50 {
	// 		return price + (price * 0.1)
	// 	}
	// 	return price
	// }

	// for product, price := range soccerProducts {
	// 	printPrice(product, price, calc)
	// }

	// //6.1
	// watersportsProducts := map[string]float64{
	// 	"Kayak":      275,
	// 	"Lifejacket": 48.95,
	// }
	// soccerProducts := map[string]float64{
	// 	"Soccer Ball": 19.50,
	// 	"Stadium":     79500,
	// }
	// waterCalc := priceCalcFactory(100, 0.2)
	// soccerCalc := priceCalcFactory(50, 0.1)
	// for product, price := range watersportsProducts {
	// 	printPrice(product, price, waterCalc)
	// }
	// for product, price := range soccerProducts {
	// 	printPrice(product, price, soccerCalc)
	// }

	// //7. Понимание оценки замыкания
	// watersportsProducts := map[string]float64{
	// 	"Kayak":      275,
	// 	"Lifejacket": 48.95,
	// }
	// soccerProducts := map[string]float64{
	// 	"Soccer Ball": 19.50,
	// 	"Stadium":     79500,
	// }
	// prizeGiveaway = false
	// waterCalc := priceCalcFactory(100, 0.2)
	// prizeGiveaway = true
	// soccerCalc := priceCalcFactory(50, 0.1)
	// // Замыкания оцениваются при вызове функции
	// for product, price := range watersportsProducts {
	// 	printPrice(product, price, waterCalc)
	// }
	// for product, price := range soccerProducts {
	// 	printPrice(product, price, soccerCalc)
	// }

	//7.1 Принудительная ранняя оценка
}
