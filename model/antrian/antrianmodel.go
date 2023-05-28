package antrian

import (
	DB "progres/db"
	entity "progres/entity/antrian"
)

func PushAntrian(data entity.Antrian) {
	if DB.JumlahAntrian != 4 {
		newData := entity.LinkedlistAntrian{}
		newData.Data = data
		temp := &DB.DBAntrian
		if temp.Next == nil {
			temp.Next = &newData
		} else {
			for temp.Next != nil {
				temp = temp.Next
			}
			temp.Next = &newData
		}
		DB.JumlahAntrian++
	} else {
		println("antrianview penuh")
	}
}

func PopAntrian() *entity.Antrian {
	if DB.DBAntrian.Next == nil {
		return nil
	}

	data := DB.DBAntrian.Next.Data

	DB.DBAntrian.Next = DB.DBAntrian.Next.Next
	if DB.DBAntrian.Next == nil {
		DB.NoAntrian = 0
	} else {
		DB.JumlahAntrian--
	}

	return &data
}

func ViewAllAntrian() []entity.Antrian {
	temp := DB.DBAntrian.Next
	var members []entity.Antrian
	for temp != nil {
		members = append(members, temp.Data)
		temp = temp.Next
	}
	return members
}

func GenerateNomorAntrian() int {
	DB.NoAntrian = DB.NoAntrian + 1
	return DB.NoAntrian
}

func IsFull() bool {
	return DB.JumlahAntrian == 4
}
