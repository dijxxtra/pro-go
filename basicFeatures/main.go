package main

import (
	"fmt"
	"sort"
)

//"math/rand"

func main() {
	// fmt.Println("Value:", rand.Int())
	// fmt.Println("Hello, Go")
	// fmt.Println(20 + 20)
	// fmt.Println(20 + 30)

	// // Типизированные константы
	// const price float32 = 275.00
	// const tax float32 = 27.5

	// // Нетепизированная константа
	// const quantity = 2
	// fmt.Println("Total:", quantity*(price+tax))

	// const price, tax float32 = 275, 27.5
	// const quantity, inStock = 2, true
	// // Литеральные значения - это нетепизированные константы
	// fmt.Println("Total:", 2*quantity*(price+tax))
	// fmt.Println("In stock:", inStock)

	// var price float32 = 275
	// var tax float32 = 27.5
	// fmt.Println(price + tax)
	// price = 300
	// fmt.Println(price + tax)

	// // Отсутствие типа не имеет такого же эффекта, как для констант
	// var price = 275.00 // Тип float64
	// var tax float32 = 27.5
	// fmt.Println(price + tax) // Ошибка

	// // Пропуск присвоения значения переменной
	// var price float32
	// fmt.Println(price)
	// price = 275.00
	// fmt.Println(price)

	// // Использование краткого синтаксиса объявления переменных
	// price := 275.00
	// fmt.Println(price)

	// price, tax, inStock := 275.00, 27.5, true
	// fmt.Println("Total:", price+tax)
	// fmt.Println("In stock:", inStock)
	// // Использование краткого синтаксиса переменных для переопределения переменных
	// price2, tax := 200.00, 25.00
	// fmt.Println("Total 2:", price2+tax)

	// // Использование пустого идентификатора
	// price, tax, inStock, _ := 275.00, 27.50, true, true
	// var _ = "Alice"
	// fmt.Println("Total:", price+tax)
	// fmt.Println("In stock:", inStock)

	// // Работа с указателями
	// first := 100
	// second := first
	// first++
	// fmt.Println("First:", first)   // First: 101
	// fmt.Println("Second:", second) // Second: 100

	// first := 100
	// var second *int = &first
	// first++
	// fmt.Println("First:", first)   // First: 101
	// fmt.Println("Second:", second) // Second: 0xc00000a0e8

	// // Следование указателю
	// first := 100
	// second := &first
	// first++
	// fmt.Println("First:", first) // First: 101
	// // *следовать указателю
	// fmt.Println("Second:", *second) // Second: 101

	// first := 100
	// second := &first
	// first++
	// *second++
	// fmt.Println("First:", first) // 102
	// fmt.Println("Second:", *second) // 102

	// first := 100
	// second := &first
	// first++
	// *second++
	// var myNewPointer *int
	// myNewPointer = second
	// *myNewPointer++
	// fmt.Println("First:", first) // First: 103
	// fmt.Println("Second:", *second) // Second: 103

	// // Нулевое значение указателей
	// first := 100
	// var second *int
	// fmt.Println(second) // nil
	// second = &first
	// fmt.Println(second) // 0xc00000a0e8

	// // Указывание на указатели
	// first := 100
	// second := &first
	// third := &second
	// fmt.Println(first)
	// fmt.Println(*second)
	// fmt.Println(**third)

	// // Почему указатели полезны
	// names := [3]string{"Alice", "Charlie", "Bob"}
	// secondName := names[1]  // Значение копируется в переменную
	// fmt.Println(secondName) // Charlie
	// sort.Strings(names[:])  // Поэтому сортировка не влияет на результат
	// fmt.Println(secondName) // Charlie

	names := [3]string{"Alice", "Charlie", "Bob"}
	secondPosition := &names[1]
	fmt.Println(*secondPosition) // Charlie
	sort.Strings(names[:])
	fmt.Println(*secondPosition) // Bob
}
