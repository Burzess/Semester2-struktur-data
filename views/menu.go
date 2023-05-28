package views

import "fmt"

func MenuAntrian() {
	fmt.Println()
	fmt.Println("===> Program Antrian Rumah Sakit <===")
	fmt.Println("1. Ambil Nomor Antrian")
	fmt.Println("2. Antrian Selanjutnya")
	fmt.Println("3. Tampilkan Semua Antrian")
	fmt.Println("4. Keluar")
}

func OlahData() {
	fmt.Println()
	fmt.Println("===> Program Olah Data Rumah Sakit <===")
	println("1. Dokter")
	println("2. Suster")
	println("3. Pasien")
	println("4. Exit")
	fmt.Print("Pilih: ")
}

func SubMenuOlahData() {
	fmt.Println("1. Insert")
	fmt.Println("2. Update")
	fmt.Println("3. Delete")
	fmt.Println("4. View By Id")
	fmt.Println("5. View All")
	fmt.Println("6. Exit")
	fmt.Print("Pilih: ")
}
