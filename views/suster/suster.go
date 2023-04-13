package viewsuster

import (
	"bufio"
	"fmt"
	"os"
)

func MenuInsert() {
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
