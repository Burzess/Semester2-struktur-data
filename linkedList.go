
package main

import (
	"fmt"
)

type Mahasiswa struct {
	nim    int
	nama   Nama
	alamat string
	next   *Mahasiswa
}

type Nama struct {
	firstName  string
	middleName string
	lastName   string
}

func insert(head *Mahasiswa, nim int, nama Nama, alamat string) *Mahasiswa {
	newMhs := &Mahasiswa{nim, nama, alamat, nil}
	if head == nil {
		return newMhs
	} else {
		temp := head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = newMhs
	}
	return head
}

func display(head *Mahasiswa) {
	if head == nil {
		fmt.Println("Linked list kosong")
	} else {
		temp := head
		for temp != nil {
			fmt.Println("addres: ", &head)
			fmt.Println("NIM:", temp.nim)
			fmt.Println("Nama:", temp.nama.firstName, temp.nama.middleName, temp.nama.lastName)
			fmt.Println("Alamat:", temp.alamat)
			fmt.Println("")
			temp = temp.next
		}
	}
}

func main() {
	var head *Mahasiswa
	var nama Nama
	fmt.Println("addres: ", &head)
	fmt.Print("Masukan nama: ")
	fmt.Scan(&nama.firstName, &nama.middleName, &nama.lastName)

	head = insert(head, 123, nama, "Jakarta")
	insert(head, 456, Nama{"Alif", "Farhan", "Cahyadi"}, "Surabaya")
	insert(head, 789, Nama{"Ahmad", "Rizal", "Efendi"}, "Bandung")

	display(head)
}
