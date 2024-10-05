package main

import "fmt"

func main() {

	// // Использование операторов if
	// kayakPrice := 275.00
	// if kayakPrice > 100 {
	// 	fmt.Println("Price is greater than 100")
	// }

	// // Использование ключевого слова else
	// kayakPrice := 275.00
	// if kayakPrice > 500 {
	// 	fmt.Println("Price is greater than 500")
	// } else if kayakPrice < 300 {
	// 	fmt.Println("Price is less than 300")
	// }

	// kayakPrice := 275.00
	// if kayakPrice > 500 {
	// 	fmt.Println("Price is greater than 500")
	// } else if kayakPrice < 100 {
	// 	fmt.Println("Price is less than 100")
	// } else if kayakPrice > 200 && kayakPrice < 300 {
	// 	fmt.Println("Price is between 200 and 300")
	//}

	// kayakPrice := 275.00
	// if kayakPrice > 500 {
	// 	fmt.Println("Price is greater than 500")
	// } else if kayakPrice < 100 {
	// 	fmt.Println("Price is less than 100")
	// } else {
	// 	fmt.Println("Price is not matched by earlier expressions")
	// }

	// // Понимание области действия оператора if
	// kayakPrice := 275.00
	// if kayakPrice > 500 {
	// 	scopedVar := 500
	// 	fmt.Println("Price is greater than", scopedVar)
	// } else if kayakPrice < 100 {
	// 	scopedVar := "Price is less than 100"
	// 	fmt.Println(scopedVar)
	// } else {
	// 	scopedVar := false
	// 	fmt.Println("Matched:", scopedVar)
	// }

	// // Использование оператора инициализации с оператором if
	// priceString := "275"
	// if kayakPrice, err := strconv.Atoi(priceString); err == nil {
	// 	fmt.Println("Price:", kayakPrice)
	// } else {
	// 	fmt.Println("Error:", err)
	//}

	// // Использование циклов for
	// counter := 0
	// for {
	// 	fmt.Println("Counter:", counter)
	// 	counter++
	// 	if counter > 3 {
	// 		break
	// 	}
	// }

	// // Включение условия в цикл
	// counter := 0
	// for counter <= 3 {
	// 	fmt.Println(counter)
	// 	counter++
	// }

	// // Использование операторов инициализации и завершения
	// // Оператор инициализации, условие и пост-оператор
	// for counter := 0; counter <= 3; counter++ {
	// 	fmt.Println(counter)
	// }

	// //Do while в Go (выполнится хотя бы один раз)
	// for counter := 10; true; counter++ {
	// 	fmt.Println("Counter:", counter)
	// 	if counter > 3 {
	// 		break
	// 	}
	// }

	// // Продолжение цикла
	// for counter := 0; counter <= 3; counter++ {
	// 	if counter == 1 {
	// 		continue
	// 	}
	// 	fmt.Println("Counter:", counter)
	// }

	// // Перечисление последовательностей
	// product := "Kayak"
	// for index, character := range product {
	// 	fmt.Println("Index:", index, "Character:", string(character))
	// }

	// // Перечисление встроенных структур данных
	// products := []string{"Kayal", "Lifejacket", "SoccerBall"}
	// for index, element := range products {
	// 	fmt.Println("Index:", index, "Element:", element)
	// }

	// // Использование операторов switch
	// product := "Kayak"
	// for index, character := range product {
	// 	switch character {
	// 	case 'K':
	// 		fmt.Println("K at position", index)
	// 	case 'y':
	// 		fmt.Println("y at position", index)
	// 	}
	// }

	// product := "Kayak"
	// for index, character := range product {
	// 	switch character {
	// 	case 'K', 'k':
	// 		fmt.Println("K or k at position", index)
	// 	case 'y':
	// 		fmt.Println("y at position", index)
	// 	}
	// }

	// product := "Kayak"
	// for index, character := range product {
	// 	switch character {
	// 	case 'K', 'k':
	// 		if character == 'k' {
	// 			fmt.Println("Lowercase k at position", index)
	// 			break
	// 		}
	// 		fmt.Println("Uppercase K at position", index)
	// 	case 'y':
	// 		fmt.Println("y at position", index)
	// 	}
	// }

	// Принудительный переход к следующему оператору case
	// product := "Kayak"
	// for index, character := range product {
	// 	switch character {
	// 	case 'K':
	// 		fmt.Println("Uppercase character")
	// 		fallthrough
	// 	case 'k':
	// 		fmt.Println("k at position", index)
	// 	case 'y':
	// 		fmt.Println("y at position", index)
	// 	}
	// }

	// // Предоставление пункта по умолчанию
	// product := "Kayak"
	// for index, character := range product {
	// 	switch character {
	// 	case 'K', 'k':
	// 		if character == 'k' {
	// 			fmt.Println("Lowercase k at position", index)
	// 			break
	// 		}
	// 		fmt.Println("Uppercase K at position", index)
	// 	case 'y':
	// 		fmt.Println("y at position", index)
	// 	default:
	// 		fmt.Println("Character", string(character), "at position", index)
	// 	}
	// }

	// // Использование оператора инициализации
	// for counter := 0; counter < 20; counter++ {
	// 	switch counter / 2 {
	// 	case 2, 3, 5, 7:
	// 		fmt.Println("Prime value:", counter/2)
	// 	default:
	// 		fmt.Println("Non-prime value:", counter/2)
	// 	}
	// }

	// for counter := 0; counter < 20; counter++ {
	// 	switch val := counter / 2; val {
	// 	case 2, 3, 5, 7:
	// 		fmt.Println("Prime value:", val)
	// 	default:
	// 		fmt.Println("Non-prime value:", val)
	// 	}
	// }

	// // Исключение значения сравнения
	// for counter := 0; counter < 10; counter++ {
	// 	switch {
	// 	case counter == 0:
	// 		fmt.Println("Zero value")
	// 	case counter < 3:
	// 		fmt.Println(counter, "is < 3")
	// 	case counter >= 3 && counter < 7:
	// 		fmt.Println(counter, "is >=3 and < 7")
	// 	default:
	// 		fmt.Println(counter, "is >= 7")
	// 	}
	// }

	// Использование операторов меток
	counter := 0
target:
	fmt.Println("Counter:", counter)
	counter++
	if counter < 5 {

		goto target
	}
}
