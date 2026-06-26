package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Domino struct {
	Kiri  int
	Kanan int
	Balak bool
}

type Dominoes struct {
	Kartu [28]Domino
	Sisa  int
}

// Membuat satu set kartu domino
func initDomino(d *Dominoes) {
	idx := 0
	for i := 0; i <= 6; i++ {
		for j := i; j <= 6; j++ {
			d.Kartu[idx] = Domino{i, j, i == j}
			idx++
		}
	}
	d.Sisa = 28
}

// Mengocok kartu
func kocokKartu(d *Dominoes) {
	rand.Seed(time.Now().UnixNano())

	for i := d.Sisa - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.Kartu[i], d.Kartu[j] = d.Kartu[j], d.Kartu[i]
	}
}

// Mengambil kartu paling atas
func ambilKartu(d *Dominoes) Domino {
	d.Sisa--
	return d.Kartu[d.Sisa]
}

// Mengambil gambar kiri atau kanan
func gambarKartu(k Domino, suit int) int {
	if suit == 0 {
		return k.Kiri
	}
	return k.Kanan
}

// Nilai kartu
func nilaiKartu(k Domino) int {
	return k.Kiri + k.Kanan
}

func main() {
	var deck Dominoes

	initDomino(&deck)
	kocokKartu(&deck)

	kartu := ambilKartu(&deck)

	fmt.Println("Kartu:", kartu.Kiri, "|", kartu.Kanan)
	fmt.Println("Nilai:", nilaiKartu(kartu))
	fmt.Println("Balak:", kartu.Balak)
}