package main

import (
	"fmt"
)

func main() {
	// // Работа с массивами
	// var names [3]string
	// names[0] = "Kayak"
	// names[1] = "Lifejacket"
	// names[2] = "Paddle"
	// fmt.Println(names)

	// // Использование литерального синтаксиса массива
	// names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	// fmt.Println(names)

	// // СОЗДАНИЕ МНОГОМЕРНЫХ МАССИВОВ
	// var coords [3][3]int
	// coords[1][2] = 10
	// fmt.Println(coords) // [[0 0 0] [0 0 10] [0 0 0]]

	// // Понимание типов массивов
	// names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	// var otherArray [4]string = names // Ошибка т.к. names имеет тип [3]string, а otherArray [4]string! Это разные типы

	// // ПОЗВОЛЯЕМ КОМПИЛЯТОРУ ОПРЕДЕЛЯТЬ ДЛИНУ МАССИВА
	// names := [...]string{"Kayak", "Lifejacket", "Paddle"}

	// // Понимание значений массива
	// names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	// otherArray := names // Копирует массив и содержащиеся в нем значения
	// names[0] = "Canoe"
	// fmt.Println("Names:", names)
	// fmt.Println("OtherArray:", otherArray)

	// names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	// otherArray := &names // *[3]string
	// names[0] = "Canoe"
	// fmt.Println("Names:", names)
	// fmt.Println("OtherArray:", *otherArray)

	// // Сравнение массивов
	// names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	// moreNames := [3]string{"Kayak", "Lifejacket", "Paddle"}

	// same := names == moreNames
	// fmt.Println("comparison:", same)

	// // Перечисление массива
	// names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	// for index, value := range names {
	// 	fmt.Println("Index:", index, "Value:", value)
	// }

	// names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	// for _, value := range names {
	// 	fmt.Println("Value:", value)
	// }

	// // Работа со срезами
	// names := make([]string, 3)
	// names[0] = "Kayak"
	// names[1] = "Lifejacket"
	// names[2] = "Paddle"
	// fmt.Println(names)

	// // Создание срезов с использованием литерального синтаксиса
	// names := []string{"Kayak", "Lifejacket", "Paddle"}
	// fmt.Println(names)

	// // Добавление элементов в срез
	// names := []string{"Kayak", "Lifejacket", "Paddle"}
	// names = append(names, "Hat", "Gloves")
	// fmt.Println(names)

	// names := []string{"Kayak", "Lifejacket", "Paddle"}
	// appendNames := append(names, "Hat", "Gloves")
	// names[0] = "Canoe"
	// fmt.Println("names:", names)                                      //Names: [Canoe Lifejacket Paddle]
	// fmt.Println("len:", len(names), "| cap:", cap(names))             // len: 3 | cap: 3
	// fmt.Println("appendNames:", appendNames)                          // appendNames: [Kayak Lifejacket Paddle Hat Gloves]
	// fmt.Println("len:", len(appendNames), "| cap:", cap(appendNames)) // len: 5 | cap: 6

	// // Выделение дополнительной емкости срезов
	// names := make([]string, 3, 6)
	// names[0] = "Kayak"
	// names[1] = "Lifejacket"
	// names[2] = "Paddle"
	// fmt.Println("len:", len(names), "| cap:", cap(names)) // len: 3 | cap: 6

	// // Функция append вызывается для среза с достаточной емкостью
	// names := make([]string, 3, 6)
	// names[0] = "Kayak"
	// names[1] = "Lifejacket"
	// names[2] = "Paddle"
	// appendNames := append(names, "Hat", "Gloves")
	// names[0] = "Canoe"
	// fmt.Println(names)       // [Canoe Lifejacket Paddle]
	// fmt.Println(appendNames) // [Canoe Lifejacket Paddle Hat Gloves]

	// // Добавление одного среза к другому
	// names := make([]string, 3, 6)
	// names[0] = "Kayak"
	// names[1] = "Lifejacket"
	// names[2] = "Paddle"
	// moreNames := []string{"Hat Gloves"}
	// appendNames := append(names, moreNames...)
	// fmt.Println("appendNames:", appendNames)
	// fmt.Println(&names[0] == &appendNames[0]) // Проверка на единный базовый массив

	// // Создание срезов из существующих массивов
	// products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	// someNames := products[1:3] // "Lifejacket", "Paddle", "Hat"
	// allNames := products[:]
	// fmt.Println("someNames:", someNames) // someNames: [Lifejacket Paddle]
	// fmt.Println("allNames:", allNames)   // allNames: [Kayak Lifejacket Paddle Hat]

	// // Добавление элементов при использовании существующих массивов для срезов
	// products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	// someNames := products[1:3]
	// allNames := products[:]
	// someNames = append(someNames, "Gloves")
	// fmt.Println("someNames:", someNames) // someNames: [Lifejacket Paddle Gloves]
	// fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
	// fmt.Println("allNames:", allNames) // allNames: [Kayak Lifejacket Paddle Gloves]
	// fmt.Println("allNames len", len(allNames), "cap:", cap(allNames))

	// // Указание емкости при создании среза из массива
	// products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	// someNames := products[1:3:3]
	// allNames := products[:]
	// someNames = append(someNames, "Boots")
	// fmt.Println("someNames:", someNames)
	// fmt.Println("allNames:", allNames)

	// // Создание срезов из других срезов
	// products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	// allNames := products[1:]  // "Lifejacket", "Paddle", "Hat"
	// someNames := allNames[1:] // "Paddle", "Hat"
	// allNames = append(allNames, "Gloves")
	// allNames[0] = "Canoe"

	// fmt.Println("someNames:", someNames) // "Paddle", "Hat"
	// fmt.Println("allNames:", allNames)   // "Canoe", "Paddle", "Hat", "Cloves"

	// // Использование функции копирования
	// ///Использование функции копирования для обеспечения разделения массива срезов
	// products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	// allNames := products[1:]
	// someNames := make([]string, 2)
	// // Получатель < Источник
	// copy(someNames, allNames)

	// fmt.Println("someNames:", someNames) // someNames: [Lifejacket Paddle]
	// fmt.Println("allNames:", allNames) // allNames: [Lifejacket Paddle Hat]

	// // Понимание ловушки неинициализированных срезов
	// products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	// allNames := products[1:]
	// var someNames []string
	// copy(someNames, allNames)
	// fmt.Println("someNames:", someNames) // []
	// fmt.Println("allNames:", allNames) // [Lifejacket Paddle Hat]

	// // Указание диапазонов при копировании срезов
	// products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	// allNames := products[1:]
	// someNames := []string{"Boots", "Canoe"}
	// copy(someNames[1:], allNames[2:3])
	// fmt.Println("allNames:", allNames)   // allNames: [Lifejacket Paddle Hat]
	// fmt.Println("someNames:", someNames) // someNames: [Boots Hat]

	// // Копирование срезов разного размера
	// products := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	// replacementProducts := []string{"Canoe", "Boots"}
	// copy(products, replacementProducts)
	// fmt.Println("products:", products) // products: [Canoe Boots Paddle Hat]

	// products := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	// replacementProducts := []string{"Canoe", "Boots"}
	// copy(products[0:1], replacementProducts)
	// fmt.Println("products:", products) // products: [Canoe Lifejacket Paddle Hat]

	// // Удаление элементов среза
	// products := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	// products = append(products[:2], products[3:]...)
	// fmt.Println(products)

	// // Перечисление срезов
	// products := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	// for index, value := range products[2:] {
	// 	fmt.Println("Index:", index, "Value:", value)
	// }

	// // Сортировка срезов
	// products := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	// sort.Strings(products)
	// slices.Sort(products)
	// for index, value := range products {
	// 	fmt.Println("Index:", index, "Value:", value)
	// }

	// // Сравнение срезов
	// p1 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	// p2 := p1
	// fmt.Println("Equal:", p1 == p2) // Ошибка

	// p1 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	// p2 := p1
	// fmt.Println("Equal:", reflect.DeepEqual(p1, p2)) // true

	// // Получение массива, лежащего на основе среза
	// p1 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	// arrayPtr := (*[3]string)(p1)
	// array := *arrayPtr
	// fmt.Println(array)

	// // Работа с картами
	// products := make(map[string]float64)
	// products["Kayak"] = 279
	// products["Lifejacket"] = 48.95
	// fmt.Println("Map size:", len(products))
	// fmt.Println("Price:", products["Kayak"])
	// fmt.Println("Price:", products["Hat"])

	// // Использование литерального синтаксиса карты
	// products := map[string]float64{
	// 	"Kayak":      279,
	// 	"Lifejacket": 48.95,
	// }
	// fmt.Println(products)

	// // Проверка элементов в карте
	// products := map[string]float64{
	// 	"Kayak":      279,
	// 	"Lifejacket": 48.95,
	// 	"Hat":        0,
	// }
	// value, ok := products["Hat"]
	// if ok {
	// 	fmt.Println("Значение существует и равно:", value)
	// } else {
	// 	fmt.Println("Такого ключа нет в карте!")
	// }

	// // Упрощенный вариант
	// products := map[string]float64{
	// 	"Kayak":      279,
	// 	"Lifejacket": 48.95,
	// 	"Hat":        0,
	// }
	// if value, ok := products["Hat"]; ok {
	// 	fmt.Println("Значение существует и равно:", value)
	// } else {
	// 	fmt.Println("Такого ключа нет!")
	// }

	// // Удаление объектов с карты
	// products := map[string]float64{
	// 	"Kayak":      279,
	// 	"Lifejacket": 48.95,
	// 	"Hat":        0,
	// }
	// delete(products, "Hat")
	// if value, ok := products["Hat"]; ok {
	// 	fmt.Println("Значение существует и равно:", value)
	// } else {
	// 	fmt.Println("Такого ключа нет!")
	// }

	// // Перечисление содержимого карты
	// products := map[string]float64{
	// 	"Kayak":      279,
	// 	"Lifejacket": 48.95,
	// 	"Hat":        0,
	// }
	// for key, value := range products {
	// 	fmt.Printf("Key: %s, Value: %.2f\n", key, value)
	// }

	// // Перечисление карты по порядку
	// products := map[string]float64{
	// 	"Kayak":      279,
	// 	"Lifejacket": 49.95,
	// 	"Hat":        0,
	// }
	// keys := make([]string, 0, len(products))
	// for key := range products {
	// 	keys = append(keys, key)
	// }
	// sort.Strings(keys)
	// for _, key := range keys {
	// 	fmt.Println("Key:", key, "Value:", products[key])
	// }

	// // Понимание двойной природы строк
	// price  := "$48.95" // string
	// currency  := price[0] // byte
	// amountString  := price[1:] // string
	// amount, err := strconv.ParseFloat(amountString, 64)
	// fmt.Println("Currency:", currency) // 36
	// if err == nil {
	// 	fmt.Println("Amount:", amount) // 48.95
	// } else {
	// 	fmt.Println("ParseErr:", err)
	// }

	// // Интерпретация byte как символа
	// price := "$48.95"
	// currency := string(price[0])
	// amountString := price[1:]
	// amount, err := strconv.ParseFloat(amountString, 64)
	// fmt.Println("Currency:", currency)
	// if err == nil {
	// 	fmt.Println("Amount:", amount)
	// } else {
	// 	fmt.Println("Parse err:", err)
	// }

	// price := "€48.95" // € - 3 байта
	// currency := string(price[0])
	// amountString := price[1:]
	// fmt.Println(len(price)) // 8 байт
	// amount, err := strconv.ParseFloat(amountString, 64)
	// fmt.Println("Currency:", currency)
	// if err == nil {
	// 	fmt.Println("Amount:", amount)
	// } else {
	// 	fmt.Println("Parse err:", err)
	// }

	// // Преобразование строки в руны
	// var price []rune = []rune("€48.95")
	// var currency string = string(price[0])
	// var amountString string = string(price[1:])
	// amount, err := strconv.ParseFloat(amountString, 64)
	// fmt.Println("Len:", len(price))    // Len: 6
	// fmt.Println("Currency:", currency) // €
	// if err == nil {
	// 	fmt.Println("Amount:", amount) // Amount: 48.95
	// } else {
	// 	fmt.Println("Parse err:", err)
	// }

	// Перечисление строк
	var price string = "€48.95"

	for index, char := range price {
		fmt.Println("Index:", index, "Char:", char, "str(Char):", string(char))
	}
	// > Index: 0 Char: 8364 str(Char): €
	// > Index: 3 Char: 52 str(Char): 4
	// > Index: 4 Char: 56 str(Char): 8
	// > Index: 5 Char: 46 str(Char): .
	// > Index: 6 Char: 57 str(Char): 9
	// > Index: 7 Char: 53 str(Char): 5
}
