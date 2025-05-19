package main

import (
	"fmt"
	"strings"
)

func tampilkanStatistik(p Pengguna) {
	fmt.Println("=== Statistik Persentase Kecocokan Karier ===")

	for i := 0; i < len(dataKarier); i++ {
		jumlahCocok := 0

		// Cek kecocokan minat
		for j := 0; j < p.banyakminat; j++ {
			for k := 0; k < len(dataKarier[i].minatCocok); k++ {
				if strings.ToLower(p.minat[j]) == strings.ToLower(dataKarier[i].minatCocok[k]) {
					jumlahCocok++
					break
				}
			}
		}

		// Cek kecocokan keahlian
		for j := 0; j < p.banyakkeahlian; j++ {
			for k := 0; k < len(dataKarier[i].keahlianDibutuhkan); k++ {
				if strings.ToLower(p.keahlian[j]) == strings.ToLower(dataKarier[i].keahlianDibutuhkan[k]) {
					jumlahCocok++
					break
				}
			}
		}

		total := p.banyakminat + p.banyakkeahlian
		var persentase float64 =0
		if total > 0 {
			persentase = (float64(jumlahCocok) / float64(total)) * 100
		}

		fmt.Printf("%s - Kecocokan: %.2f%%\n", dataKarier[i].namaKarier, persentase)
	}
}
