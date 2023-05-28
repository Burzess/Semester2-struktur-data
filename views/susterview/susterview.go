package susterview

import (
	"fmt"
	controller "progres/controller/sustercontroller"
)

func InsertSuster() {
	var id int
	var nama, tlp, shift string
	fmt.Print("masukkan ID Baru : ")
	fmt.Scan(&id)
	fmt.Print("masukkan Nama Baru : ")
	fmt.Scan(&nama)
	fmt.Print("masukan nomor telpon : ")
	fmt.Scan(&tlp)
	fmt.Print("masukan shift : ")
	fmt.Scan(&shift)

	err := controller.ControllerInsertSuster(id, nama, tlp, shift)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("insert berhasil")
	}
}

func UpdateSuster() {
	var id int
	var nama, tlp, shift string
	fmt.Print("masukkan ID : ")
	fmt.Scan(&id)
	fmt.Print("masukkan Nama Baru : ")
	fmt.Scan(&nama)
	err := controller.ControllerUpdate(id, nama, tlp, shift)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("update berhasil")
	}
}

func DeleteDoter() {
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

func ViewSuster() {
	allSuster := controller.ControllerFindAllSuster()
	// fmt.Println(allSuster)
	for _, member := range allSuster {
		fmt.Println("id Suster :", member.Id)
		fmt.Println("Nama Suster :", member.Nama)
	}
}
