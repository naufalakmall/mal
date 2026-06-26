package main

import "fmt"

func main() {
	var x string
	var n int

	fmt.Print("String yang dicari: ")
	fmt.Scan(&x)

	fmt.Print("Jumlah data: ")
	fmt.Scan(&n)

	data := make([]string, n)

	fmt.Println("Masukkan data string:")

	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	ditemukan := false
	posisi := -1
	jumlah := 0

	for i := 0; i < n; i++ {
		if data[i] == x {
			jumlah++
			if !ditemukan {
				ditemukan = true
				posisi = i + 1
			}
		}
	}

	if ditemukan {
		fmt.Println("a. String ditemukan")
	} else {
		fmt.Println("a. String tidak ditemukan")
	}

	if ditemukan {
		fmt.Println("b. Posisi pertama:", posisi)
	} else {
		fmt.Println("b. Tidak ada posisi")
	}

	fmt.Println("c. Jumlah string:", jumlah)

	if jumlah >= 2 {
		fmt.Println("d. Ya, ada minimal dua")
	} else {
		fmt.Println("d. Tidak")
	}
}