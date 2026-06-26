package main

import "fmt"

type Domino struct {
	Kiri  int
	Kanan int
}

func bisaMain(meja Domino, kartu Domino) bool {

	return meja.Kiri == kartu.Kiri ||
		meja.Kiri == kartu.Kanan ||
		meja.Kanan == kartu.Kiri ||
		meja.Kanan == kartu.Kanan
}

func main() {

	meja := Domino{4, 5}

	tangan := []Domino{
		{6, 6},
		{1, 3},
		{2, 5},
		{0, 1},
		{5, 5},
	}

	fmt.Println("Kartu yang bisa dimainkan:")

	for _, k := range tangan {

		if bisaMain(meja, k) {
			fmt.Println(k.Kiri, "|", k.Kanan)
		}
	}
}