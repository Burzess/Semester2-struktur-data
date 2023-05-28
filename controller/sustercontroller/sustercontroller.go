package sustercontroller

import (
	"errors"
	entities "progres/entity/suster"
	model "progres/model/suster"
)

func ControllerInsertSuster(id int, nama, tlp, shift string) error {
	find := model.Search(id)
	if find == nil {
		data := entities.Suster{
			Id:    id,
			Nama:  nama,
			Tlp:   tlp,
			Shift: shift,
		}
		model.ModelInsertSuster(data)
		return nil
	}

	return errors.New("insert Data gagal, id yang di inputkan sudah tersedia")
}

func ControllerUpdate(id int, nama, tlp, shift string) error {
	container := entities.Suster{
		Id:    id,
		Nama:  nama,
		Tlp:   tlp,
		Shift: shift,
	}

	if model.ModelUpdateSuster(container) {
		return nil
	}

	return errors.New("update gagal id tidak ditemukan")
}

func ControllerDelete(id int) error {
	err := model.ModelDeleteSuster(id)
	if err {
		return nil
	}

	return errors.New("delete gagal id tidak ditemukan")
}

func ControllerViewById(id int) *entities.LinkedlistSuster {
	current := model.Search(id)
	if current == nil {
		return nil
	}
	return current
}

func ControllerFindAllSuster() []entities.Suster {
	return model.ModelViewAllSuster()
}
