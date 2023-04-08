package controllers

import (
	"tugas-uts/entities"
	"tugas-uts/model/dosen"
	"tugas-uts/model/mahasiswa"
)

// fungsi dosen
func InsertDataDosen(g *entities.GerbongDsn, data entities.Dosen) {
	dosen.InsertDataDosen(g, data)
}

func UpdateDataDosen(g *entities.GerbongDsn, nip string) {
	dosen.Update(g, nip)
}

func DeleteDataDosen(g *entities.GerbongDsn, nip string) {
	dosen.Delete(g, nip)
}

func GetListDataDosen(g *entities.GerbongDsn) {
	dosen.GetList(g)
}

func GetlistDataDosenByNip(g *entities.GerbongDsn, nip string) {
	dosen.GetListByNip(g, nip)
}

// fungsi dosen end

// fungsi mahasiswa
func InsertDataMahasiswa(g *entities.GerbongMhs, data entities.Mahasiswa) {
	mahasiswa.InsertDataMahasiswa(g, data)
}

func UpdateDataMahasiswa(g *entities.GerbongMhs, npm string) {
	mahasiswa.Update(g, npm)
}

func DeleteDataMahasiswa(g *entities.GerbongMhs, npm string) {
	mahasiswa.Delete(g, npm)
}

func GetListDataMahasiswa(g *entities.GerbongMhs) {
	mahasiswa.GetListMahasiswa(g)
}

func GetlistDatMahasiswaByNpm(g *entities.GerbongMhs, npm string) {
	mahasiswa.GetlistMahasiswaByNpm(g, npm)
}

//fungsi mahasiswa end
