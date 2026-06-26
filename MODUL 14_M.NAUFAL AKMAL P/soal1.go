package main

import "fmt"

const MAX = 100

func insertionSort(arr *[MAX]int, n int) {
	var temp, j int

	for i := 1; i < n; i++ {
		temp = arr[i]
		j = i - 1

		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = temp
	}
}

func jarakTetap(arr [MAX]int, n int) bool {
	if n <= 2 {
		return true
	}

	selisih := arr[1] - arr[0]

	for i := 2; i < n; i++ {
		if arr[i]-arr[i-1] != selisih {
			return false
		}
	}
	return true
}

func main() {
	var data [MAX]int
	var n int
	var x int

	fmt.Println("Masukkan bilangan (akhiri dengan bilangan negatif):")

	for {
		fmt.Scan(&x)

		if x < 0 {
			break
		}

		data[n] = x
		n++
	}

	insertionSort(&data, n)

	fmt.Println("\nData setelah diurutkan:")
	for i := 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	if jarakTetap(data, n) {
		fmt.Println("Data berjarak")
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}