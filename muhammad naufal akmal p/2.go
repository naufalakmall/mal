package main

import "fmt"

func main() {
	var Hahut int

	fmt.Print("Hahut: ")
	fmt.Scanln(&Hahut)

	if Hahut%400 == 0 || (Hahut%4 == 0 && Hahut%100 != 0) {
		fmt.Println("Kabisat: true")
	} else {
		fmt.Println("Kabisat: false")
	}
}