package main

import "fmt"

var teks string
var idx int

func start() {
	idx = 0
}

func maju() {
	idx++
}

func eop() bool {
	return idx >= len(teks) || teks[idx] == '.'
}

func cc() byte {
	return teks[idx]
}

func main() {

	fmt.Print("Masukkan kalimat (akhiri titik): ")
	fmt.Scanln(&teks)

	start()

	for !eop() {
		fmt.Printf("%c ", cc())
		maju()
	}
}