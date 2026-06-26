package main

import "fmt"

func main() {
	var x, jumlah float64
	var banyak int

	fmt.Println("Masukkan bilangan (9999 untuk berhenti):")

	for {
		fmt.Scan(&x)

		if x == 9999 {
			break
		}

		jumlah += x
		banyak++
	}

	if banyak > 0 {
		fmt.Printf("Rata-rata = %.2f\n", jumlah/float64(banyak))
	} else {
		fmt.Println("Tidak ada data.")
	}
}