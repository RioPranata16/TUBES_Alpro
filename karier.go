package main
import "fmt"

type Karier struct{
	namaKarier string
	kategori string
	gaji float64
	minatCocok [5]string
	keahlianDibutuhkan [5]string
	jumlahMinatCocok int
	jumlahKeahlianDibutuhkan int
}
//list data karier
var dataKarier [10]Karier
func inisialisasiKarier(){
	dataKarier [0] =Karier{
		namaKarier: "Data Scientist",
		kategori: "Teknologi",
		gaji: 12000000,
		minatCocok: [5]string{"Analisis Data","Pemrograman","","",""},
		keahlianDibutuhkan: [5]string{"Python", "Statistik","","",""},
		jumlahMinatCocok: 2,
		jumlahKeahlianDibutuhkan: 2,
	}
	dataKarier [1]=Karier{
		namaKarier: "UI/UX Designer",
		kategori: "Desain",
		gaji: 9000000,
		minatCocok: [5]string{"Desain","Psikologi Pengguna","","",""},
		keahlianDibutuhkan: [5]string{"Figma", "Desain Grafis","","",""},
		jumlahMinatCocok: 2,
		jumlahKeahlianDibutuhkan: 2,
	}
	dataKarier [2]=Karier{
		namaKarier: "Digital Marketer",
		kategori: "Pemasaran",
		gaji: 8000000,
		minatCocok: [5]string{"Strategi","Komunikasi","","",""},
		keahlianDibutuhkan: [5]string{"SEO", "Social Media","","",""},
		jumlahMinatCocok: 2,
		jumlahKeahlianDibutuhkan: 2,
	}
	dataKarier [3]=Karier{
		namaKarier: "Software Engineer",
		kategori: "Teknologi",
		gaji: 13000000,
		minatCocok: [5]string{"Pemrograman","Problem Solving","","",""},
		keahlianDibutuhkan: [5]string{"Golang", "Git","","",""},
		jumlahMinatCocok: 2,
		jumlahKeahlianDibutuhkan: 2,
	}
	dataKarier [4]=Karier{
		namaKarier: "Data Analyst",
		kategori: "Teknologi",
		gaji: 10000000,
		minatCocok: [5]string{"Data","Statistik","","",""},
		keahlianDibutuhkan: [5]string{"SQL", "Excel","","",""},
		jumlahMinatCocok: 2,
		jumlahKeahlianDibutuhkan: 2,
	}
	dataKarier [5]=Karier{
		namaKarier: "Content Writer",
		kategori: "Kreatif",
		gaji: 7000000,
		minatCocok: [5]string{"Menulis","Kreativitas","","",""},
		keahlianDibutuhkan: [5]string{"Copywriting", "Riset","","",""},
		jumlahMinatCocok: 2,
		jumlahKeahlianDibutuhkan: 2,
	}
	dataKarier [6]=Karier{
		namaKarier: "Product Manager",
		kategori: "Manajemen Produk",
		gaji: 14000000,
		minatCocok: [5]string{"Organisasi","Inovasi","","",""},
		keahlianDibutuhkan: [5]string{"Agile", "Komunikasi","","",""},
		jumlahMinatCocok: 2,
		jumlahKeahlianDibutuhkan: 2,
	}
	dataKarier [7]=Karier{
		namaKarier: "Financial Analyst",
		kategori: "Keuangan",
		gaji: 11000000,
		minatCocok: [5]string{"Angka","Ekonomi","","",""},
		keahlianDibutuhkan: [5]string{"Excel", "Akutansi","","",""},
		jumlahMinatCocok: 2,
		jumlahKeahlianDibutuhkan: 2,
	}
	dataKarier [8]=Karier{
		namaKarier: "Network Engineer",
		kategori: "Infrastruktur",
		gaji: 10500000,
		minatCocok: [5]string{"Jaringan","Teknologi","","",""},
		keahlianDibutuhkan: [5]string{"Cisco", "Troubleshooting","","",""},
		jumlahMinatCocok: 2,
		jumlahKeahlianDibutuhkan: 2,
	}
	dataKarier [9]=Karier{
		namaKarier: "HR Specialist",
		kategori: "Sumber Daya Manusia",
		gaji: 8500000,
		minatCocok: [5]string{"Sosial","Administrasi","","",""},
		keahlianDibutuhkan: [5]string{"rekrutmen", "Undang-Undang Ketenagakerjaan","","",""},
		jumlahMinatCocok: 2,
		jumlahKeahlianDibutuhkan: 2,
	}
}
//fungsi untuk merekomendasikan karier berdasarkan minat dan keahlian
func cocokKarier(p Pengguna) []Rekomendasi {
	var rekomendasi []Rekomendasi

	for i := 0; i < len(dataKarier); i++ {
		var skor int = 0

		// Cek kecocokan minat
		for j := 0; j < p.banyakminat; j++ {
			for k := 0; k < dataKarier[i].jumlahMinatCocok; k++ {
				if p.minat[j] == dataKarier[i].minatCocok[k] {
					skor++
					break
				}
			}
		}

		// Cek kecocokan keahlian
		for j := 0; j < p.banyakkeahlian; j++ {
			for k := 0; k < dataKarier[i].jumlahKeahlianDibutuhkan; k++ {
				if p.keahlian[j] == dataKarier[i].keahlianDibutuhkan[k] {
					skor++
					break
				}
			}
		}

		// Tambahkan hasil ke slice rekomendasi
		rekomendasi = append(rekomendasi, Rekomendasi{
			namaKarier: dataKarier[i].namaKarier,
			skor:       skor,
			gaji:       dataKarier[i].gaji,
		})
	}
	return rekomendasi
}

//fungsi untuk menambah minat dan keahlian
func tambahMinatKeahlian(index int) {
    var jumlahMinatBaru int
    fmt.Print("Masukkan jumlah minat yang ingin ditambahkan: ")
    fmt.Scanln(&jumlahMinatBaru)

    // Hitung total minat setelah penambahan
    totalMinat := dataKarier[index].jumlahMinatCocok + jumlahMinatBaru
    if totalMinat > 5 {
        fmt.Printf("Maksimal 5 minat saja. Kamu sudah punya %d minat.\n", dataKarier[index].jumlahMinatCocok)
        return
    }

    for i := 0; i < jumlahMinatBaru; i++ {
        fmt.Printf("Masukkan minat ke-%d: ", dataKarier[index].jumlahMinatCocok+i+1)
        fmt.Scanln(&dataKarier[index].minatCocok[dataKarier[index].jumlahMinatCocok+i])
    }
    dataKarier[index].jumlahMinatCocok = totalMinat

    var jumlahKeahlianBaru int
    fmt.Print("Masukkan jumlah keahlian yang ingin ditambahkan: ")
    fmt.Scanln(&jumlahKeahlianBaru)

    totalKeahlian := dataKarier[index].jumlahKeahlianDibutuhkan + jumlahKeahlianBaru
    if totalKeahlian > 5 {
        fmt.Printf("Maksimal 5 keahlian saja. Kamu sudah punya %d keahlian.\n", dataKarier[index].jumlahKeahlianDibutuhkan)
        return
    }

    for i := 0; i < jumlahKeahlianBaru; i++ {
        fmt.Printf("Masukkan keahlian ke-%d: ", dataKarier[index].jumlahKeahlianDibutuhkan+i+1)
        fmt.Scanln(&dataKarier[index].keahlianDibutuhkan[dataKarier[index].jumlahKeahlianDibutuhkan+i])
    }
    dataKarier[index].jumlahKeahlianDibutuhkan = totalKeahlian
}

//fungsi untuk ubah minat dan keahlian
func ubahMinatKeahlian(index int) {
	var jumlahMinat int
	fmt.Print("Masukkan jumlah minat yang cocok: ")
	fmt.Scanln(&jumlahMinat)

	// Update minat berdasarkan jumlah yang diinput
	for i := 0; i < jumlahMinat; i++ {
		fmt.Print("Masukkan minat ke-", i+1, ": ")
		fmt.Scanln(&dataKarier[index].minatCocok[i])
	}
	dataKarier[index].jumlahMinatCocok = jumlahMinat

	// Update jumlah keahlian yang dibutuhkan
	var jumlahKeahlian int
	fmt.Print("Masukkan jumlah keahlian yang dibutuhkan: ")
	fmt.Scanln(&jumlahKeahlian)

	// Update keahlian berdasarkan jumlah yang diinput
	for i := 0; i < jumlahKeahlian; i++ {
		fmt.Print("Masukkan keahlian ke-", i+1, ": ")
		fmt.Scanln(&dataKarier[index].keahlianDibutuhkan[i])
	}
	dataKarier[index].jumlahKeahlianDibutuhkan = jumlahKeahlian
}
// Fungsi untuk menghapus minat atau keahlian
func hapusMinatKeahlian(index int) {
	var jenis string
	fmt.Print("Ingin menghapus minat atau keahlian? (minat/keahlian): ")
	fmt.Scanln(&jenis)

	if jenis == "minat" {
		var minatHapus int
		fmt.Print("Masukkan nomor minat yang ingin dihapus (1 sampai ", dataKarier[index].jumlahMinatCocok, "): ")
		fmt.Scanln(&minatHapus)

		if minatHapus >= 1 && minatHapus <= dataKarier[index].jumlahMinatCocok {
			// Hapus minat dengan mengganti nilai yang dihapus menjadi string kosong
			dataKarier[index].minatCocok[minatHapus-1] = ""
			dataKarier[index].jumlahMinatCocok--
			fmt.Println("Minat berhasil dihapus.")
		} else {
			fmt.Println("Nomor minat tidak valid.")
		}
	} else if jenis == "keahlian" {
		var keahlianHapus int
		fmt.Print("Masukkan nomor keahlian yang ingin dihapus (1 sampai ", dataKarier[index].jumlahKeahlianDibutuhkan, "): ")
		fmt.Scanln(&keahlianHapus)

		if keahlianHapus >= 1 && keahlianHapus <= dataKarier[index].jumlahKeahlianDibutuhkan {
			// Hapus keahlian dengan mengganti nilai yang dihapus menjadi string kosong
			dataKarier[index].keahlianDibutuhkan[keahlianHapus-1] = ""
			dataKarier[index].jumlahKeahlianDibutuhkan--
			fmt.Println("Keahlian berhasil dihapus.")
		} else {
			fmt.Println("Nomor keahlian tidak valid.")
		}
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func tampilanKarier(){
	fmt.Println("===Daftar Karier===")
	for i :=0;i<len(dataKarier);i++{
		if dataKarier[i].namaKarier !=""{
			fmt.Println("Nama Karier:", dataKarier[i].namaKarier)
			fmt.Println("Kategori:",dataKarier[i].kategori)
			fmt.Printf("Gaji: %.0f\n", dataKarier[i].gaji)
			fmt.Println("Minat Cocok: ", dataKarier[i].minatCocok)
			fmt.Println("Keahlian Dibutuhkan: ", dataKarier[i].keahlianDibutuhkan)
		}
	}
}
