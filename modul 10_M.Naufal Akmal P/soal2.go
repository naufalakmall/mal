package main

import "fmt"

func main() {
	var x, y int
	var berat [1000]float64
	var totalWadah [1000]float64

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	jumlahWadah := x / y
	if x%y != 0 {
		jumlahWadah++
	}

	for i := 0; i < x; i++ {
		indexWadah := i / y
		totalWadah[indexWadah] += berat[i]
	}

	var totalSemua float64

	for i := 0; i < jumlahWadah; i++ {
		fmt.Print(totalWadah[i])
		if i < jumlahWadah-1 {
			fmt.Print(" ")
		}
		totalSemua += totalWadah[i]
	}

	fmt.Println()

	rataRata := totalSemua / float64(jumlahWadah)
	fmt.Println(rataRata)
}