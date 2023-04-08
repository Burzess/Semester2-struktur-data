package dosenview

import (
	"fmt"
	"tugas-uts/controllers"
	"tugas-uts/entities"
)

func MenuInsertDosen(dataDsn *entities.GerbongDsn) {
	fmt.Println("anda masuk pilihan 1: INSERT")
	var nip string
	var nama string
	fmt.Print("masukkan nip: ")
	fmt.Scan(&nip)
	fmt.Print("masukkan nama: ")
	fmt.Scan(&nama)

	data := entities.Dosen{ // data yang di inputkan di simpan di data struct Dosen yang ada di model
		Nip:  nip,
		Nama: nama,
	}
	controllers.InsertDataDosen(dataDsn, data) // memanggil fungsi Insert() di controllers
}

func MenuUpdateDosen(dataDsn *entities.GerbongDsn) {
	fmt.Println("anda masuk pilihan 2: UPDATE")
	var nip string
	fmt.Println("Masukan ID yang ingin di Update : ")
	fmt.Scan(&nip)
	controllers.UpdateDataDosen(dataDsn, nip)
}

func MenuDeleteDosen(dataDsn *entities.GerbongDsn) {
	fmt.Println("anda masuk pilihan 3: DELETE")
	var nip string
	fmt.Print("Masukan ID yang ingin delete : ")
	fmt.Scan(&nip)
	controllers.DeleteDataDosen(dataDsn, nip)
}

func MenuViewAll(dataDsn *entities.GerbongDsn) {
	fmt.Println("anda masuk pilihan 4: VIEW ALL")
	controllers.GetListDataDosen(dataDsn)
}

func MenuViewByNip(dataDsn *entities.GerbongDsn) {
	fmt.Println("anda masuk pilihan 5: VIEW BY NIP")
	var nip string
	fmt.Print("masukkan NIP : ")
	fmt.Scan(&nip)
	controllers.GetlistDataDosenByNip(dataDsn, nip)
}
