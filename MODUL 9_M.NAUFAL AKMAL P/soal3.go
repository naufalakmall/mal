package main

import "fmt"

const MAX = 100

func main() {
	var klubA, klubB string
	var pemenang [MAX]string
	var skorA, skorB int
	var jumlah int
	var pertandingan int = 1

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)

	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	for {
		fmt.Printf("Pertandingan %d : ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang[jumlah] = klubA
			jumlah++
		} else if skorB > skorA {
			pemenang[jumlah] = klubB
			jumlah++
		} else {
			pemenang[jumlah] = "Draw"
			jumlah++
		}

		pertandingan++
	}

	fmt.Println("\nHasil Pertandingan")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("Hasil %d : %s\n", i+1, pemenang[i])
	}

	fmt.Println("Pertandingan selesai")
}