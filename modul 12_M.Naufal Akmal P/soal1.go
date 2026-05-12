package main

import "fmt"

func main() {
	var suara int
	var suaraMasuk int
	var suaraSah int

	hitung := make([]int, 21)

	for {
		fmt.Scan(&suara)

		if suara == 0 {
			break
		}

		suaraMasuk++

		if suara >= 1 && suara <= 20 {
			hitung[suara]++
			suaraSah++
		}
	}

	fmt.Println("Suara masuk:", suaraMasuk)
	fmt.Println("Suara sah:", suaraSah)

	for i := 1; i <= 20; i++ {
		if hitung[i] > 0 {
			fmt.Printf("%d: %d\n", i, hitung[i])
		}
	}
}