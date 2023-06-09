package suster

import (
	DB "progres/db"
	entity "progres/entity/suster"
)

func ModelInsertSuster(container entity.Suster) {
	newGerbong := entity.LinkedlistSuster{}
	newGerbong.Data = container
	temp := &DB.DBSuster
	if temp.Next == nil {
		temp.Next = &newGerbong
	} else {
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &newGerbong
	}
}

func Search(id int) *entity.LinkedlistSuster {
	current := DB.DBSuster.Next
	for current != nil {
		if current.Data.Id == id {
			return current
		}
		current = current.Next
	}
	return nil
}

func ModelUpdateSuster(newData entity.Suster) bool {
	current := &DB.DBSuster
	for current != nil {
		if current.Data.Id == newData.Id {
			current.Data = newData
			return true
		}
		current = current.Next
	}
	return false
}

func ModelDeleteSuster(id int) bool {
	prev := &DB.DBSuster
	current := DB.DBSuster.Next

	for current != nil {
		if current.Data.Id == id {
			prev.Next = current.Next
			return true
		}
		prev = current
		current = current.Next
	}

	return false
}

func ModelViewAllSuster() []entity.Suster {
	temp := DB.DBSuster.Next
	members := []entity.Suster{}
	for temp != nil {
		members = append(members, temp.Data)
		temp = temp.Next
	}
	return members
}
