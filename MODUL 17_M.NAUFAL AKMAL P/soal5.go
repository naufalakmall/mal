package main

import (
	"fmt"
	"math"
)

func main() {

	var n int
	fmt.Print("N suku pertama: ")
	fmt.Scan(&n)

	var jumlah float64
	var sebelumnya float64

	for i := 0; i < n; i++ {

		suku := 1.0 / float64(2*i+1)

		if i%2 == 0 {
			jumlah += suku
		} else {
			jumlah -= suku
		}

		if i > 0 {
			if math.Abs(jumlah-sebelumnya) <= 0.00001 {
				fmt.Printf("Hasil PI : %.9f\n", jumlah*4)
				fmt.Printf("Pada i ke : %d\n", i+1)
				return
			}
		}

		sebelumnya = jumlah
	}

	fmt.Printf("Hasil PI : %.9f\n", jumlah*4)
}