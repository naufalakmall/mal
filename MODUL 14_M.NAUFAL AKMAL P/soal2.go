package main

import (
	"fmt"
)

const nMax = 7919

type Buku struct {
	ID        string
	Judul     string
	Penulis   string
	Penerbit  string
	Eksemplar int
	Tahun     int
	Rating    int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(n)

	fmt.Println("Masukkan data buku:")
	for i := 0; i < *n; i++ {
		fmt.Println("\nData Buku", i+1)

		fmt.Print("ID        : ")
		fmt.Scan(&pustaka[i].ID)

		fmt.Print("Judul     : ")
		fmt.Scan(&pustaka[i].Judul)

		fmt.Print("Penulis   : ")
		fmt.Scan(&pustaka[i].Penulis)

		fmt.Print("Penerbit  : ")
		fmt.Scan(&pustaka[i].Penerbit)

		fmt.Print("Eksemplar : ")
		fmt.Scan(&pustaka[i].Eksemplar)

		fmt.Print("Tahun     : ")
		fmt.Scan(&pustaka[i].Tahun)

		fmt.Print("Rating    : ")
		fmt.Scan(&pustaka[i].Rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		return
	}

	max := pustaka[0].Rating
	idx := 0

	for i := 1; i < n; i++ {
		if pustaka[i].Rating > max {
			max = pustaka[i].Rating
			idx = i
		}
	}

	fmt.Println("\n=== Buku Terfavorit ===")
	fmt.Println("Judul    :", pustaka[idx].Judul)
	fmt.Println("Penulis  :", pustaka[idx].Penulis)
	fmt.Println("Penerbit :", pustaka[idx].Penerbit)
	fmt.Println("Tahun    :", pustaka[idx].Tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var temp Buku
	var j int

	for i := 1; i < n; i++ {
		temp = pustaka[i]
		j = i - 1

		for j >= 0 && pustaka[j].Rating < temp.Rating {
			pustaka[j+1] = pustaka[j]
			j--
		}

		pustaka[j+1] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	fmt.Println("\n=== 5 Buku Rating Tertinggi ===")

	batas := 5
	if n < 5 {
		batas = n
	}

	for i := 0; i < batas; i++ {
		fmt.Printf("%d. %s (Rating %d)\n",
			i+1,
			pustaka[i].Judul,
			pustaka[i].Rating)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	kiri := 0
	kanan := n - 1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2

		if pustaka[tengah].Rating == r {
			fmt.Println("\n=== Buku Ditemukan ===")
			fmt.Println("Judul     :", pustaka[tengah].Judul)
			fmt.Println("Penulis   :", pustaka[tengah].Penulis)
			fmt.Println("Penerbit  :", pustaka[tengah].Penerbit)
			fmt.Println("Tahun     :", pustaka[tengah].Tahun)
			fmt.Println("Eksemplar :", pustaka[tengah].Eksemplar)
			fmt.Println("Rating    :", pustaka[tengah].Rating)
			return
		}

		if r > pustaka[tengah].Rating {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	fmt.Println("\nTidak ada buku dengan rating seperti itu")
}

func main() {
	var pustaka DaftarBuku
	var n int
	var ratingCari int

	DaftarkanBuku(&pustaka, &n)

	CetakTerfavorit(pustaka, n)

	UrutBuku(&pustaka, n)

	Cetak5Terbaru(pustaka, n)

	fmt.Print("\nMasukkan rating yang dicari: ")
	fmt.Scan(&ratingCari)

	CariBuku(pustaka, n, ratingCari)
}