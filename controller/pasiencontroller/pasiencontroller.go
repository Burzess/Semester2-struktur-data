package pasiencontroller

import (
	"errors"
	entities "progres/entity/pasien"
	model "progres/model/pasien"
)

func ControllerInsertPasien(id int, nama, tlp, penyakit string) error {
	find := model.Search(id)
	if find == nil {
		data := entities.Pasien{
			Id:       id,
			Nama:     nama,
			Telpon:   tlp,
			Penyakit: penyakit,
		}
		model.ModelInsertPasien(data)
		return nil
	}

	return errors.New("insert Data gagal, id yang di inputkan sudah tersedia")
}

func ControllerUpdate(id int, nama, tlp, penyakit string) error {
	container := entities.Pasien{
		Id:       id,
		Nama:     nama,
		Telpon:   tlp,
		Penyakit: penyakit,
	}

	if model.ModelUpdatePasien(container) {
		return nil
	}

	return errors.New("update gagal id tidak ditemukan")
}

func ControllerDelete(id int) error {
	err := model.ModelDeletePasien(id)
	if err {
		return nil
	}

	return errors.New("delete gagal id tidak ditemukan")
}

func ControllerViewById(id int) *entities.LinkedlistPasien {
	current := model.Search(id)
	if current == nil {
		return nil
	}
	return current
}

func ControllerFindAllPasien() []entities.Pasien {
	return model.ModelViewAllPasien()
}
