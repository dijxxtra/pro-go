package main

// Данная глава описывает команды go (fmt, build,install, run, vet и т.д), отладчик и линтеры.

import "fmt"

func main() {
	fmt.Println("Hello, Go")
	for i:=0; i<5; i++ {
		fmt.Println(i)
	}
}
