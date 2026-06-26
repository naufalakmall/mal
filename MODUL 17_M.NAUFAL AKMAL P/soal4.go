package main

import "fmt"

func main() {

	var n int

	fmt.Print("N suku pertama: ")
	fmt.Scan(&n)

	var jumlah float64

	for i := 0; i < n; i++ {

		suku := 1.0 / float64(2*i+1)

		if i%2 == 0 {
			jumlah += suku
		} else {
			jumlah -= suku
		}
	}

	fmt.Printf("Hasil PI: %.7f\n", jumlah*4)
}