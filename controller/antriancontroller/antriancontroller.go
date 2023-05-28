package antriancontroller

import (
	entities "progres/entity/antrian"
	"progres/model/antrian"
	"progres/model/dokter"
)

func PushAntrian(pilih int, nama, keluhan string) {
	noAntrian := antrian.GenerateNomorAntrian()
	dokter := dokter.PilihDokter(pilih)

	data := entities.Antrian{
		NamaDokter: dokter,
		NoAntrian:  noAntrian,
		Name:       nama,
		Keluhan:    keluhan,
	}

	antrian.PushAntrian(data)
}

func ViewAllAntrian() []entities.Antrian {
	return antrian.ViewAllAntrian()
}

func PopAntrian() *entities.Antrian {
	return antrian.PopAntrian()
}

func IsFull() bool {
	return antrian.IsFull()
}
