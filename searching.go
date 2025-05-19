package main
import "fmt"
import "strings"

func cariKarierbyNama(nama string)int{
	for i :=0;i<len(dataKarier);i++{
		if dataKarier[i].namaKarier==nama{
			return i
		}
	}
	return -1
}
func karierygDicari(idx int){
	fmt.Println("===Karier Yang Anda Cari===")
	if idx >=0 && idx <len(dataKarier){
		fmt.Println("Nama Karier:", dataKarier[idx].namaKarier)
		fmt.Println("Kategori:", dataKarier[idx].kategori)
		fmt.Println("Gaji:", dataKarier[idx].gaji)
		fmt.Println("Minat Cocok:", dataKarier[idx].minatCocok)
		fmt.Println("Keahlian yang dibutuhkan:", dataKarier[idx].keahlianDibutuhkan)
	}else{
		fmt.Println("Karier Tidak Ditemukan")
	}
}
func cariKarierbyNamaInsensitive(nama string) int {
    for i := 0; i < len(dataKarier); i++ {
        if strings.ToLower(dataKarier[i].namaKarier) == nama {
            return i
        }
    }
    return -1
}
