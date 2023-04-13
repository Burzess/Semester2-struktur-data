package dokterview

import (
	"bufio"
	"fmt"
	"os"
	"tugas-uts/entities"
)

func MenuInsert(dataDokter *entities.Dokter) {
	scanner := bufio.NewScanner(os.Stdin)
	var id int
	var nama string
	var Tlp string
	var jamKerja string
	var spesialis string
	fmt.Print("Masukan ID : ")
	fmt.Scan(&id)
	scanner.Scan()
	fmt.Print("Masukan Nama : ")
	fmt.Scan(&nama)
	scanner.Scan()
	fmt.Print("Masukan Telpon : ")
	fmt.Scan(&Tlp)
	scanner.Scan()
	fmt.Print("Masukan Jam Kerja : ")
	fmt.Scan(&jamKerja)
	scanner.Scan()
	fmt.Print("Masukan Spesialis : ")
	fmt.Scan(&spesialis)
	scanner.Scan()
	doktercontrollers.InsertDokter(dataDokter, id, nama, Tlp, jamKerja, spesialis)
}

func MenuUpdate() {
	scanner := bufio.NewScanner(os.Stdin)
	var id int
	var nama string
	var Tlp string
	var jamKerja string
	var spesialis string
	fmt.Print("Masukan ID : ")
	fmt.Scan(&id)
	scanner.Scan()
	fmt.Print("Masukan Nama : ")
	fmt.Scan(&nama)
	scanner.Scan()
	fmt.Print("Masukan Telpon : ")
	fmt.Scan(&Tlp)
	scanner.Scan()
	fmt.Print("Masukan Jam Kerja : ")
	fmt.Scan(&jamKerja)
	scanner.Scan()
	fmt.Print("Masukan Spesialis : ")
	fmt.Scan(&spesialis)
	scanner.Scan()
}

func MenuDelete() {
	var id int
	fmt.Print("Masukan ID")
	fmt.Scan(&id)
}

func MenuViewById() {
	var id int
	fmt.Print("Masukan ID")
	fmt.Scan(&id)
}
