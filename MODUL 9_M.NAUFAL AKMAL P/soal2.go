package main

import (
	"fmt"
	"math"
)

const MAX = 100

func main() {
	var arr [MAX]int
	var n int

	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("\na. Seluruh isi array:")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Println("\nb. Elemen pada indeks ganjil:")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Println("\nc. Elemen pada indeks genap:")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Print("\nMasukkan nilai x: ")
	fmt.Scan(&x)

	fmt.Println("d. Elemen pada indeks kelipatan", x, ":")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var hapus int
	fmt.Print("\nMasukkan indeks yang akan dihapus: ")
	fmt.Scan(&hapus)

	for i := hapus; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	n--

	fmt.Println("e. Array setelah penghapusan:")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var jumlah int
	for i := 0; i < n; i++ {
		jumlah += arr[i]
	}
	rata := float64(jumlah) / float64(n)
	fmt.Printf("\nf. Rata-rata = %.2f\n", rata)

	var total float64
	for i := 0; i < n; i++ {
		selisih := float64(arr[i]) - rata
		total += selisih * selisih
	}
	stdDev := math.Sqrt(total / float64(n))
	fmt.Printf("g. Standar Deviasi = %.2f\n", stdDev)

	var cari, frek int
	fmt.Print("\nMasukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&cari)

	for i := 0; i < n; i++ {
		if arr[i] == cari {
			frek++
		}
	}

	fmt.Printf("h. Frekuensi %d = %d\n", cari, frek)
}