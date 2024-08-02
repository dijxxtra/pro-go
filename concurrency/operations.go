package main

import (
	"fmt"
	"time"
)

func CalcStoreTotal(data ProductData) {
	var storeTotal float64
	var channel chan float64 = make(chan float64, 2)

	for category, group := range data {
		go group.TotalPrice(category, channel)
	}

	time.Sleep(time.Second * 5)
	fmt.Println("--Начинает получать с канала")

	for i := 0; i < len(data); i++ {
		fmt.Println("--Канал ожидает прочтения", len(channel), "элементов в буфере, размер:", cap(channel))
		value := <-channel
		fmt.Println("--канал прочитал значение", value)
		storeTotal += value
		time.Sleep(time.Second)
	}
	fmt.Println("Итого:", ToCurrency(storeTotal))
}

func (group ProductGroup) TotalPrice(category string, resultChannel chan float64) {
	var total float64
	for _, p := range group {
		total += p.Price
	}
	fmt.Println(category, "канал отправил ", ToCurrency(total))
	resultChannel <- total
	fmt.Println(category, "канал отправка завершена")
}
