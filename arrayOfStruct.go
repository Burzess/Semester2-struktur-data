package main

import "fmt"

type Nama struct {
	firstName  string
	middleName string
	lastName   string
}

type Mahasiswa struct {
	npm  string
	nama Nama
	ipk  float64
}

var dbMahasiswa = [10]Mahasiswa{}
var index int = 0

func main() {
	mahasiswa2 := Mahasiswa{
		npm:  "06.2022.1.07619",
		nama: Nama{"Mas", "Ibrahim", "Halim"},
		ipk:  4.00,
	}

	mahasiswa3 := Mahasiswa{
		npm:  "06.2022.1.07621",
		nama: Nama{"Alif", "Farhan", "Cahyadi"},
		ipk:  3.89,
	}

	insertDbMahasiswa(mahasiswa2)
	insertDbMahasiswa(mahasiswa3)

	display()
}
