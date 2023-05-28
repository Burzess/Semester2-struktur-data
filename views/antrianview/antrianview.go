package antrianview

import (
	"fmt"
	"progres/controller/antriancontroller"
	"progres/controller/doktercontroller"
)

func TambahAntrian() {
	if !antriancontroller.IsFull() {
		var (
			pilih         int
			nama, Keluhan string
		)
		fmt.Print("Masukan Nama Antrian: ")
		fmt.Scan(&nama)
		fmt.Print("Masukan Keluhan: ")
		fmt.Scan(&Keluhan)
		fmt.Println("List Dokter")
		dokter := doktercontroller.ControllerFindAllDokter()
		count := 1
		for _, current := range dokter {
			fmt.Printf("  %d. %s\n", count, current.Nama)
			count++
		}
		fmt.Print("Pilih Dokter: ")
		fmt.Scan(&pilih)
		antriancontroller.PushAntrian(pilih, nama, Keluhan)
		fmt.Println("Nomor antrian baru telah ditambahkan")
		fmt.Println()
	} else {
		fmt.Println("\n==>Antrian saat ini sedang penuh<==")
		fmt.Println()
	}
}

func NextAntrian() {
	currAntrian := antriancontroller.PopAntrian()
	if currAntrian != nil {
		println()
		println("\t\tPANGGILAN ANTRIAN")
		fmt.Printf("NO ANTRIAN : %d\n", currAntrian.NoAntrian)
		fmt.Printf("ATAS NAMA  : %s\n", currAntrian.Name)
		fmt.Printf("Silahkan menemui dokter %s di ruangannya\n\n", currAntrian.NamaDokter)
	} else {
		fmt.Println("\n==>ANTRIAN KOSONG<==")
	}
}

func ViewAllAntrian() {
	currAntrian := antriancontroller.ViewAllAntrian()
	if currAntrian != nil {
		println()
		for _, Antrian := range currAntrian {
			fmt.Printf("NO ANTRIAN : %d\n", Antrian.NoAntrian)
			fmt.Printf("ATAS NAMA  : %s\n\n", Antrian.Name)
		}
	} else {
		fmt.Println("\n==>ANTRIAN KOSONG<==")
	}

}
