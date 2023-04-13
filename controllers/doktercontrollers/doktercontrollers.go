package doktercontrollers

import (
	"tugas-uts/Tugas_UTS/doktermodel"
)

func InsertDokter(dataDokter *entities.Dokter, id int, nama string, Tlp string, jamKerja string, spesialis string) {
	doktermodel.InsertDokter(dataDokter, id, nama, Tlp, jamKerja, spesialis)
}
