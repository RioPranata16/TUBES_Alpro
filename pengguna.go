package main
import "fmt"

type Pengguna struct{
	minat [10] string
	banyakminat int
	keahlian [10] string
	banyakkeahlian int
}

func inputminat(p *Pengguna){
	var minatSekarang string
	for{
		fmt.Println("Masukkan Minat Anda atau ketik selesai untuk berhenti memasukan minat: ")
		fmt.Scanln(&minatSekarang)
		if minatSekarang=="selesai"{
			break
		}else if p.banyakminat <10{
			p.minat[p.banyakminat]=minatSekarang
			p.banyakminat++
		}else{
			fmt.Println("Maaf Anda Hanya bisa menginputkan minat dibawah 10 saja")
			break
		}
		
	}
}
func inputkeahlian(p *Pengguna){
	var keahlianSekarang string
	for{
		fmt.Println("Masukkan Keahlian anda atau ketik selesai untuk berhenti memasukkan keahlian: ")
		fmt.Scanln(&keahlianSekarang)
		if keahlianSekarang== "selesai"{
			break
		}else if p.banyakkeahlian <10{
			p.keahlian[p.banyakkeahlian]=keahlianSekarang
			p.banyakkeahlian++
		}else{
			fmt.Println("Maaf anda Hanya bisa menginputkan minat dibawah 10 saja")
			break
		}
	}
}

func tampilpengguna(p Pengguna){
	fmt.Println("====Daftar Minat Anda===")
	for i :=0;i<p.banyakminat;i++{
		fmt.Println("Minat ke ",i+1, ":",p.minat[i])
	}
	fmt.Println("===Daftar Keahlian Anda===")
	for i :=0;i<p.banyakkeahlian;i++{
		fmt.Println("Keahlian Ke", i+1,":", p.keahlian[i])
	}
}