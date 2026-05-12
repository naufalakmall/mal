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

	ketua := 0
	wakil := 0

	for i := 1; i <= 20; i++ {
		if hitung[i] > hitung[ketua] {
			wakil = ketua
			ketua = i
		} else if i != ketua && hitung[i] > hitung[wakil] {
			wakil = i
		}
	}

	fmt.Println("Suara masuk:", suaraMasuk)
	fmt.Println("Suara sah:", suaraSah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}