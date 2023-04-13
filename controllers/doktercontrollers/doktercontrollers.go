package doktercontrollers

import (
	"tugas-uts/doktermodel"
	"tugas-uts/entities"
)

func InsertDokter(dataDokter *entities.Dokter, id int, nama string, Tlp string, jamKerja string, spesialis string) {
	doktermodel.InsertDokter(dataDokter, id, nama, Tlp, jamKerja, spesialis)
}
