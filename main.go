package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)


func main() {
	inisialisasiKarier()
	var pilihan int
	var p Pengguna
	for {
		fmt.Println("\n=== Menu Karier ===")
		fmt.Println("[1] Tampilkan Daftar Karier")
		fmt.Println("[2] Tambah Minat & Keahlian pada Karier")
		fmt.Println("[3] Ubah Minat & Keahlian pada Karier")
		fmt.Println("[4] Hapus Minat atau Keahlian pada Karier")
		fmt.Println("[5] Rekomendasi Karier untuk Pengguna")
		fmt.Println("[6] Cari Karier Berdasarkan Nama")
		fmt.Println("[7] Urutkan Rekomendasi Karier")
		fmt.Println("[8] Statistik Presentase Kecocokan Karier")
		fmt.Println("[9] keluar")
		fmt.Print("Pilih menu (1-9): ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tampilanKarier()
		case 2:
			// inputminat(&p)
			// inputkeahlian(&p)

			var index int
			fmt.Print("Masukkan indeks karier (0-9) yang ingin ditambah minat/keahliannya: ")
			fmt.Scanln(&index)
			if index >= 0 && index < len(dataKarier) {
				tambahMinatKeahlian(index)
			} else {
				fmt.Println("Indeks tidak valid.")
			}
		case 3:
			var index int
			fmt.Print("Masukkan indeks karier (0-9) yang ingin diubah minat/keahliannya: ")
			fmt.Scanln(&index)
			if index >= 0 && index < len(dataKarier) {
				ubahMinatKeahlian(index)
			} else {
				fmt.Println("Indeks tidak valid.")
			}
		case 4:
			var index int
			fmt.Print("Masukkan indeks karier (0-9) yang ingin dihapus minat/keahliannya: ")
			fmt.Scanln(&index)
			if index >= 0 && index < len(dataKarier) {
				hapusMinatKeahlian(index)
			} else {
				fmt.Println("Indeks tidak valid.")
			}
		case 5:
				var p Pengguna
				fmt.Print("Masukkan jumlah minat Anda (maks 5): ")
				fmt.Scanln(&p.banyakminat)
				for i := 0; i < p.banyakminat; i++ {
					fmt.Print("Minat ke-", i+1, ": ")
					fmt.Scanln(&p.minat[i])
				}
				fmt.Print("Masukkan jumlah keahlian Anda (maks 5): ")
				fmt.Scanln(&p.banyakkeahlian)
				for i := 0; i < p.banyakkeahlian; i++ {
					fmt.Print("Keahlian ke-", i+1, ": ")
					fmt.Scanln(&p.keahlian[i])
				}

				// Ambil rekomendasi berdasarkan input pengguna
				rekom := cocokKarier(p)

				// Tampilkan hanya hasil yang memiliki skor > 0 (agar relevan)
				var hasilCocok []Rekomendasi
				for _, r := range rekom {
					if r.skor > 0 {
						hasilCocok = append(hasilCocok, r)
					}
				}

				// Tampilkan hasilnya
				if len(hasilCocok) == 0 {
					fmt.Println("Tidak ada karier yang cocok dengan minat dan keahlian Anda.")
				} else {
					printRekomendasiTanpaSkor(hasilCocok) // atau printRekomendasi(hasilCocok) kalau mau tampilkan skor
				}

		case 6:
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Masukkan Karier yang ingin dicari: ")
			karier, _ := reader.ReadString('\n')
			karier = strings.TrimSpace(karier)        // hapus spasi/enter di awal dan akhir
			karier = strings.ToLower(karier)          // buat lowercase semua
			// Buat fungsi cariKarierbyNama case-insensitive
			index := cariKarierbyNamaInsensitive(karier)
			if index >= 0 && index < len(dataKarier) {
				karierygDicari(index)
			} else {
				fmt.Println("Karier Tidak ditemukan")
			}
		case 7:
			var p Pengguna
			fmt.Print("Masukkan jumlah minat Anda (maks 5): ")
			fmt.Scanln(&p.banyakminat)
			for i := 0; i < p.banyakminat; i++ {
				fmt.Print("Minat ke-", i+1, ": ")
				fmt.Scanln(&p.minat[i])
			}
			fmt.Print("Masukkan jumlah keahlian Anda (maks 5): ")
			fmt.Scanln(&p.banyakkeahlian)
			for i := 0; i < p.banyakkeahlian; i++ {
				fmt.Print("Keahlian ke-", i+1, ": ")
				fmt.Scanln(&p.keahlian[i])
			}

			// Hitung rekomendasi dulu
			rekomendasi := cocokKarier(p)

			// Pilih metode sorting
			var metode, kriteria int
			fmt.Println("\nPilih metode pengurutan:")
			fmt.Println("[1] Selection Sort")
			fmt.Println("[2] Insertion Sort")
			fmt.Print("Pilihan: ")
			fmt.Scanln(&metode)

			fmt.Println("Urutkan berdasarkan:")
			fmt.Println("[1] Skor kecocokan")
			fmt.Println("[2] Gaji rata-rata")
			fmt.Print("Pilihan: ")
			fmt.Scanln(&kriteria)

			if metode == 1 && kriteria == 1 {
				selectionSortBySkor(rekomendasi)
			} else if metode == 1 && kriteria == 2 {
				selectionSortByGaji(rekomendasi)
			} else if metode == 2 && kriteria == 1 {
				insertionSortBySkor(rekomendasi)
			} else if metode == 2 && kriteria == 2 {
				insertionSortByGaji(rekomendasi)
			} else {
				fmt.Println("Pilihan metode atau kriteria tidak valid.")
				break
			}

			// Tampilkan hasil pengurutan
			printRekomendasi(rekomendasi)

		case 8:
			inputminat(&p)
			inputkeahlian(&p)
			tampilpengguna(p)
			if p.banyakminat == 0 && p.banyakkeahlian == 0 {
			fmt.Println("❗ Anda belum mengisi minat dan keahlian")
			} else {
			tampilkanStatistik(p)
			}
		case 9:
			fmt.Println("Terima kasih telah menggunakan program ini.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
