package main

import (
	"fmt"
	"strconv"
)

func main() {
	// // Использование арифметических операторов
	// price, tax := 275.00, 27.4
	// sum := price + tax
	// difference := price - tax
	// product := price * tax
	// quotient := price / tax
	// fmt.Println("Sum:", sum)
	// fmt.Println("Difference:", difference)
	// fmt.Println("Product:",product)
	// fmt.Println("Quotient:",quotient)

	// //Понимание арифметического переполнения
	// var intVal = math.MaxInt64
	// var floatVal = math.MaxFloat64
	// fmt.Println(intVal * 2)   // -2
	// fmt.Println(floatVal * 2) // +Inf
	// fmt.Println(math.IsInf((floatVal * 2), 0))

	// // Остаток от деления
	// posResult := 3 % 2
	// negResult := -3 % 2
	// absResult := math.Abs(float64(negResult))
	// fmt.Println(posResult) // 1
	// fmt.Println(negResult) // -1
	// fmt.Println(absResult) // 1

	// Понимание арифметического переполнения
	// value := 10.2
	// value++
	// fmt.Println(value) // 11.2
	// value += 2
	// fmt.Println(value) // 13.2
	// value -= 2
	// fmt.Println(value) // 11.2
	// value--
	// fmt.Println(value) // 10.2

	// // Объединение строк
	// greeting := "Hello"
	// language := "Go"
	// combinedString := greeting + ", " + language
	// fmt.Println(combinedString)

	// // Понимание операторов сравнения
	// first := 100
	// const second = 200.0
	// equal := first == second
	// notEqual := first != second
	// lessThan := first < second
	// lessThanOrEqual := first <= second
	// greaterThan := first > second
	// greaterThanOrEqual := first >= second
	// fmt.Println(equal)
	// fmt.Println(notEqual)
	// fmt.Println(lessThan)
	// fmt.Println(lessThanOrEqual)
	// fmt.Println(greaterThan)
	// fmt.Println(greaterThanOrEqual)

	// // Сравнение указателей
	// first := 100
	// second := &first
	// third := &first
	// alpha := 100
	// beta := &alpha
	// fmt.Println(second == third)
	// fmt.Println(second == beta)

	// // Понимание логических операторов
	// maxMph := 50
	// passengerCapacity := 4
	// airbags := true
	// familyCar := passengerCapacity > 2 && airbags
	// sportsCar := maxMph >= 100 || passengerCapacity == 2
	// canCategorize := !familyCar && !sportsCar
	// fmt.Println(familyCar)
	// fmt.Println(sportsCar)
	// fmt.Println(canCategorize)

	// // Преобразование, анализ и форматирование значений
	// kayak := 275
	// soccerBall := 19.5
	// total := kayak + soccerBall // Ошибка
	// fmt.Println(total)

	// // Выполнение явных преобразований
	// kayak := 275
	// soccerBall := 19.5
	// total := float64(kayak) + soccerBall
	// fmt.Println(total)

	// // Понимание ограничений явных преобразований
	// kayak := 275
	// soccerBall := 19.5
	// total := kayak + int(soccerBall) // 19.5 > 19
	// fmt.Println(total)               // 294
	// fmt.Println(int8(total))         // 38 т.к переменная переполнена

	// kayak := 275
	// soccerBall := 19.5
	// total := kayak + int(math.Round(soccerBall))
	// fmt.Println(total) // 295

	// weight := 103.4
	// fmt.Println("Weight Ceill:", math.Ceil(weight))
	// fmt.Println("Weight Floor:", math.Floor(weight))

	// // Парсинг из строк
	// val1 := "true"
	// val2 := "false"
	// val3 := "not true"
	// bool1, b1err := strconv.ParseBool(val1)
	// bool2, b2err := strconv.ParseBool(val2)
	// bool3, b3err := strconv.ParseBool(val3)
	// fmt.Println("Bool 1:", bool1, "err: ", b1err) // Bool 1: true err:  <nil>
	// fmt.Println("Bool 2:", bool2, "err: ", b2err) // Bool 2: false err:  <nil>
	// fmt.Println("Bool 3:", bool3, "err: ", b3err) // Bool 3: false err:  strconv.ParseBool: parsing "not true": invalid syntax

	// // Проверка на наличие ошибок
	// val1 := "0"
	// bool1, err := strconv.ParseBool(val1)
	// if err == nil {
	// 	fmt.Println("Parsed value:", bool1)
	// } else {
	// 	fmt.Println("Cannot parse:", err)
	// }

	// // Инициализация и проверка в операторе if
	// val1 := "TRUE"

	// if bool1, err := strconv.ParseBool(val1); err == nil {
	// 	fmt.Println("Parsed value:", bool1)
	// } else {
	// 	fmt.Println("Cannot parse:", err)
	// }

	// // Разбор целых чисел
	// val1 := "127"
	// // строка, основания для числа или ноль, размер типа данных
	// int1, err := strconv.ParseInt(val1, 0, 8)
	// if err == nil {
	// 	fmt.Println("Parsed value:", int1)
	// } else {
	// 	fmt.Println("Cannot parse:", val1)
	// }

	// val1 := "100"
	// if int1, err := strconv.ParseInt(val1, 0, 8); err == nil {
	// 	smallInt := int8(int1)
	// 	fmt.Println("Parsed value:", smallInt)
	// } else {
	// 	fmt.Println("Cannot Parse:", val1, err)
	// }

	// // Разбор двоичных, восьмеричных и шестнадцатеричных целых чисел
	// val1 := "100" // в двоичной системе счисления
	// if int8base2, err := strconv.ParseInt(val1, 2, 8); err == nil {
	// 	smallInt := int8(int8base2) // Результат всегда в 10й системе счисления
	// 	fmt.Println("Parsed value:", smallInt) // 4
	// } else {
	// 	fmt.Println("Cannot parsed:", val1, err)
	// }

	// val1 := "0b1100100"
	// if int8base2, err := strconv.ParseInt(val1, 0, 8); err == nil {
	// 	smallInt := int8(int8base2)
	// 	fmt.Println("Parsed value:", smallInt) // 100
	// } else {
	// 	fmt.Println("Cannot parse:", val1, err)
	// }

	// // Использование удобной целочисленной функции
	// val1 := "100"
	// if int1, err := strconv.ParseInt(val1, 10, 0); err == nil {
	// 	intResult := int(int1)
	// 	fmt.Println("Parsed value:", intResult)
	// } else {
	// 	fmt.Println("Cannot parse:", val1, err)
	// }

	// // Atoi
	// val1 := "100"
	// // Atoi эквивален ParseInt(val1, 10, 0)
	// if int1, err := strconv.Atoi(val1); err == nil {
	// 	fmt.Printf("Parse value: %d, type: %T\n", int1, int1)
	// } else {
	// 	fmt.Println("Cannot parse:", val1, err)
	// }

	// // Разбор чисел с плавающей запятой
	// val1 := "48.95"
	// // ParseFloat преобразует результат в float64
	// float1, err := strconv.ParseFloat(val1, 64)
	// if err == nil {
	// 	fmt.Println("Parsed value:", float1)
	// } else {
	// 	fmt.Println("Cannot Parse:", val1, err)
	// }

	// // ParseFloat работа с экспонентой
	// val1 := "4.895e+01"
	// float1, err := strconv.ParseFloat(val1, 64)
	// if err == nil {
	// 	fmt.Println("Parse value:", float1)
	// } else {
	// 	fmt.Println("Cannot Parse:", val1, err)
	// }

	// // Форматирование значений как строк
	// // Форматирование логических значений
	// val1 := true
	// val2 := false
	// str1 := strconv.FormatBool(val1)
	// str2 := strconv.FormatBool(val2)
	// fmt.Println("Formatted value 1:", str1)
	// fmt.Println("Formatted value 2:", str2)

	// // Форматирование целочисленных значений
	// val := 24
	// base10String := strconv.FormatInt(int64(val), 10)
	// base2String := strconv.FormatInt(int64(val), 2)
	// fmt.Println("Base 10:", base10String)
	// fmt.Println("Base 2:", base2String)

	// // Использование удобной целочисленной функции
	// age := 24
	// ageString := strconv.Itoa(age)
	// fmt.Println("Ваау ты преобразовал свой возраст в строку! Теперь мы знаем что тебе: " + ageString + "!")

	// // Форматирование значений с плавающей запятой
	// val := 49.95
	// Fstring := strconv.FormatFloat(val, 'f', 2, 64)
	// Estring := strconv.FormatFloat(val, 'e', -1, 64)
	// fmt.Println("Fstring:", Fstring)
	// fmt.Println("Estring:", Estring)

	// ipv4 := [4]int64{192, 168, 0, 1}
	// var ipv4Binary string
	// for index, address := range ipv4 {
	// 	if index == 3 {
	// 		ipv4Binary += fmt.Sprintf("%08s", strconv.FormatInt(address, 2))
	// 	} else {
	// 		ipv4Binary += fmt.Sprintf("%08s", strconv.FormatInt(address, 2)) + "."
	// 	}
	// }
	// fmt.Println(ipv4Binary)

}
