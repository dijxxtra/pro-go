package main

import "fmt"

func enumerateProducts(channel1, channel2 chan<- *Product) {
	for _, p := range ProductList {
		select {
		case channel1 <- p:
		case channel2 <- p:
		}
	}
	close(channel1)
	close(channel2)
}

func main() {
	prodChan1 := make(chan *Product, 2)
	prodChan2 := make(chan *Product, 2)
	go enumerateProducts(prodChan1, prodChan2)

	for p := range prodChan1 {
		fmt.Println("Channel 1 received product:", p.Name)
	}
	for p := range prodChan2 {
		fmt.Println("Channel 2 received product:", p.Name)
	}

}
