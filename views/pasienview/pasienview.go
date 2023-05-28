package pasienview

import (
	"fmt"
	controller "progres/controller/pasiencontroller"
)

func InsertPasien() {
	var id int
	var nama, tlp, penyakit string
	fmt.Print("masukkan ID : ")
	fmt.Scan(&id)
	fmt.Print("masukkan Nama : ")
	fmt.Scan(&nama)
	fmt.Print("masukan Penyakit yang di derita: ")
	fmt.Scan(&penyakit)
	fmt.Println("masukan nomor telpon: ")
	fmt.Scan(&tlp)

	err := controller.ControllerInsertPasien(id, nama, tlp, penyakit)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("insert berhasil")
	}
}

func UpdatePasien() {
	var id int
	var nama, tlp, penyakit string
	fmt.Print("masukkan ID : ")
	fmt.Scan(&id)
	fmt.Print("masukkan Nama Baru : ")
	fmt.Scan(&nama)
	fmt.Print("masukan Penyakit yang di derita: ")
	fmt.Scan(&penyakit)
	fmt.Println("masukan nomor telpon: ")
	fmt.Scan(&tlp)

	err := controller.ControllerUpdate(id, nama, tlp, penyakit)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("update berhasil")
	}
}

func DeletePasien() {
	var id int
	fmt.Println("masukan ID")
	fmt.Scan(&id)
	err := controller.ControllerDelete(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("delete berhasil")
	}
}

func ViewById() {
	var id int
	fmt.Println("masukan ID")
	fmt.Scan(&id)
	current := controller.ControllerViewById(id)
	if current != nil {
		fmt.Println("Id: ", current.Data.Id)
		fmt.Println("Nama: ", current.Data.Nama)
	} else {
		fmt.Println("data tidak di temukan")
	}
}

func ViewPasien() {
	allPasien := controller.ControllerFindAllPasien()
	for _, member := range allPasien {
		fmt.Println("id Pasien   : ", member.Id)
		fmt.Println("Nama Pasien : ", member.Nama)
		fmt.Println("Penyakit    : ", member.Penyakit)
		fmt.Println("Telpon      : ", member.Telpon)
	}
}
