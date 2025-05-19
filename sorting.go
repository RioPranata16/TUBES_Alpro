package main

import "fmt"

type Rekomendasi struct {
	namaKarier string
	skor       int
	gaji       float64
}

// === SELECTION SORT ===
func selectionSortBySkor(data []Rekomendasi) {
	for i := 0; i < len(data)-1; i++ {
		idxMax := i
		for j := i + 1; j < len(data); j++ {
			if data[j].skor > data[idxMax].skor {
				idxMax = j
			}
		}
		data[i], data[idxMax] = data[idxMax], data[i]
	}
}

func selectionSortByGaji(data []Rekomendasi) {
	for i := 0; i < len(data)-1; i++ {
		idxMax := i
		for j := i + 1; j < len(data); j++ {
			if data[j].gaji > data[idxMax].gaji {
				idxMax = j
			}
		}
		data[i], data[idxMax] = data[idxMax], data[i]
	}
}

// === INSERTION SORT ===
func insertionSortBySkor(data []Rekomendasi) {
	for i := 1; i < len(data); i++ {
		temp := data[i]
		j := i - 1
		for j >= 0 && data[j].skor < temp.skor {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
}

func insertionSortByGaji(data []Rekomendasi) {
	for i := 1; i < len(data); i++ {
		temp := data[i]
		j := i - 1
		for j >= 0 && data[j].gaji < temp.gaji {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
}

// === CETAK REKOMENDASI ===
func printRekomendasi(data []Rekomendasi) {
	fmt.Println("=== Daftar Rekomendasi Karier ===")
	for _, r := range data {
		fmt.Printf("%s - Skor: %d, Gaji: %.0f\n", r.namaKarier, r.skor, r.gaji)
	}
}
func printRekomendasiTanpaSkor(data []Rekomendasi) {
	fmt.Println("=== Daftar Rekomendasi Karier ===")
	for _, r := range data {
		fmt.Printf("%s - Gaji: %.0f\n", r.namaKarier, r.gaji)
	}
}
