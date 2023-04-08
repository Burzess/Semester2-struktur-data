package dosen

import (
	"fmt"
	"os"
	"tugas-uts/entities"

	"github.com/olekukonko/tablewriter"
)

func InsertDataDosen(g *entities.GerbongDsn, data entities.Dosen) {
	newGerbong := &entities.GerbongDsn{Dosen: data}

	temp := g
	if temp.Next == nil {
		temp.Next = newGerbong
	} else {
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = newGerbong
	}
}

func Update(g *entities.GerbongDsn, nip string) {
	temp := g

	if temp.Next == nil {
		fmt.Println("Data Kosong")
		return
	} else {
		for temp.Next != nil {
			if temp.Next.Dosen.Nip == nip {
				fmt.Print("Masukan nama baru : ")
				fmt.Scan(&temp.Next.Dosen.Nama)
				return
			}
			temp = temp.Next
		}
	}

	fmt.Printf("data dengan nip %s  tidak ada\n", nip)
}

func Delete(g *entities.GerbongDsn, nip string) {
	temp := g

	if temp.Next == nil {
		fmt.Println("Data Kosong")
		return
	}
	// mengecek apakah data yang di hapus data pertama
	if temp.Dosen.Nip == nip {
		*temp = *temp.Next
		return
	}

	for temp.Next != nil {
		if temp.Next.Dosen.Nip == nip {
			temp.Next = temp.Next.Next
			return
		}
		temp = temp.Next
	}

	fmt.Println("Delete gagal, ID tidak tersedia")
}

// Fungsi untuk mendapatkan daftar NIP dan Nama dosen pada suatu gerbong
func GetList(g *entities.GerbongDsn) {
	// Inisialisasi tabel dan header kolom
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"NIP", "Nama"})

	// Iterasi melalui setiap elemen GerbongMhs dan menambahkan data ke dalam tabel
	current := g
	for current.Next != nil {
		dsn := current.Next.Dosen
		row := []string{dsn.Nip, dsn.Nama}
		table.Append(row)
		current = current.Next
	}

	// Render tabel ke terminal
	table.Render()
}

func GetListByNip(g *entities.GerbongDsn, nip string) {
	current := g

	for current.Next != nil {
		if current.Next.Dosen.Nip == nip {
			dsn := current.Next.Dosen
			// Inisialisasi tabel dan header kolom
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"NIP", "Nama"})
			row := []string{dsn.Nip, dsn.Nama}
			table.Append(row)
			// Render tabel ke terminal
			table.Render()
			return
		}
		current = current.Next
	}
	fmt.Printf("Data dengan NIP %s tidak tersedia", nip)
}
